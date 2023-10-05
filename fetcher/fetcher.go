package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//  rateLimiter 爬取等待
	<-rateLimiter

	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	// Douban requires parameters User-Agent otherwise it will return 418 error
	req.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
	)
	req.Header.Add(
		"Accept",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	reader := getReaderEncoding(resp.Body)
	return reader, nil
}

// 获取解析为utf8编码后的 Reader
func getReaderEncoding(r io.Reader) []byte {
	reader := bufio.NewReader(r)
	e := determinEncoding(reader)
	// Convert encoding
	utf8reader := transform.NewReader(reader, e.NewDecoder())
	bytes, err := io.ReadAll(utf8reader)
	if err != nil {
		panic(err)
	}
	return bytes
}

func determinEncoding(reader *bufio.Reader) encoding.Encoding {
	// 该方式读取缓冲区1024字节而不移动指针位置
	bytes, err := reader.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	// DetermineEncoding determines the encoding of an HTML document by examining up
	// to the first 1024 bytes of content and the declared Content-Type.
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

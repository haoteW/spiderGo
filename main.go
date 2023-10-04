package main

import (
	"bufio"
	"net/http"
)

func main() {
	url := "https://movie.douban.com"
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	// Douban requires parameters User-Agent otherwise it will return 418 error
	req.Header.Add("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		// ReadString returns err != nil if and only if the returned data does not end in delim.
		line, err := reader.ReadString('\n')
		if err != nil { //读取错误或者结尾
			println(err.Error())
			break
		}
		println(line)
	}
}

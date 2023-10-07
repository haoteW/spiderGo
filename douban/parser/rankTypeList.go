package parser

import (
	"fmt"
	"good-spider/engine"
	"good-spider/model"
	"log"
	"net/url"
	"regexp"
)

const typeRegex = `<a href="(/typerank\?type_name=[^&]*?&type=[^&]*?&interval_id=[^&]*?&action=)">([^<]+?)</a>`

// 解析排行榜类别
func ParseTypeList(contents []byte, noOne model.RankType) engine.ParseResult {
	// MustCompile will panic if cannot parser
	re := regexp.MustCompile(typeRegex)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        buildRequest(string(m[1])),
				ParserFunc: ParseRankType,
			},
		)
		fmt.Printf("Type: %s, URL : %s\n", m[2], m[1])
	}
	fmt.Printf("matches found : %d\n", len(matches))
	return result
}

func buildRequest(suffix string) string {
	parsedURL, err := url.Parse(model.Domain + suffix)
	if err != nil {
		log.Print("URL parsing failed:", err)
		return ""
	}
	// 获取查询参数
	queryParams := parsedURL.Query()

	// 获取特定参数的值
	movieType := queryParams.Get("type")
	intervalID := queryParams.Get("interval_id")
	action := queryParams.Get("action")

	// 拼接新链接
	newSuffix := fmt.Sprintf("/j/chart/top_list?type=%s&interval_id=%s&action=%s&start=0&limit=20",
		movieType, intervalID, action)

	return model.Domain + newSuffix
}

package parser

import (
	"encoding/json"
	"fmt"
	"good-spider/engine"
	"good-spider/model"
)

func ParseRankType(content []byte) engine.ParseResult {
	result := engine.ParseResult{}
	// 解析接口返回结果
	var rankType model.RankType
	err := json.Unmarshal(content, &rankType)
	if err != nil {
		return result
	}
	for _, item := range rankType {
		result.Items = append(result.Items, item.Title)
		result.Requests = append(
			result.Requests, engine.Request{
				Url: item.URL,
				ParserFunc: func(contents []byte) engine.ParseResult {
					return ParseMovieInfo(contents, item.Title)
				},
			},
		)
		fmt.Printf("Type: %s, URL : %s\n", item.Title, item.URL)
	}

	return result
}

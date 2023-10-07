package engine

import "good-spider/model"

// 请求信息 url及对应句柄函数
type Request struct {
	Url        string
	ParserFunc func([]byte, model.RankType) ParseResult
	RankType   model.RankType
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

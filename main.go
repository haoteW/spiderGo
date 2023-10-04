package main

import (
	"good-spider/douban/parser"
	"good-spider/engine"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://movie.douban.com/chart",
		ParserFunc: parser.ParseTypeList,
	})
}

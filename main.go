package main

import (
	"good-spider/douban/parser"
	"good-spider/engine"
	"good-spider/schedule"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &schedule.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "https://movie.douban.com/chart",
		ParserFunc: parser.ParseTypeList,
	})
}

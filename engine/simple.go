package engine

import (
	"good-spider/fetcher"
	"log"
)

type SimpleEngine struct{}

/*
单任务版
请求队列插入、worker调度
*/
func (e *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher Error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}

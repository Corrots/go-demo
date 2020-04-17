package engine

import (
	"fmt"
	"log"

	"github.com/corrots/go-demo/crawler/fetcher"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		result, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			fmt.Printf("got item %v\n", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	fmt.Printf("fetching URL: %s\n", r.URL)
	bytes, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("fetcher err, fetching %s, %v\n", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(bytes), nil
}

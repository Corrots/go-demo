package engine

import (
	"fmt"
	"log"

	"github.com/corrots/go-demo/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		fmt.Printf("fetching URL: %s\n", r.URL)
		bytes, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("fetcher err, fetching %s, %v\n", r.URL, err)
			continue
		}
		result := r.ParserFunc(bytes)
		requests = append(requests, result.Requests...)

		for _, item := range result.Items {
			fmt.Printf("got item %v\n", item)
		}
	}
}

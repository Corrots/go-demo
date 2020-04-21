package engine

import (
	"fmt"
	"log"

	"github.com/corrots/go-demo/crawler/fetcher"
)

func worker(r Request) (ParseResult, error) {
	fmt.Printf("fetching URL: %s\n", r.URL)
	bytes, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("fetcher err, fetching %s, %v\n", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(bytes, r.URL), nil
}

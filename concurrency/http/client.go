package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

const URL = "http://www.imooc.com"

func main() {
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect: ", req)
			return nil
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("%s\n", bytes[:500])
}

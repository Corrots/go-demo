package main

import (
	"github.com/corrots/go-demo/crawler/engine"
	"github.com/corrots/go-demo/crawler/zhenai/parser"
)

const URL = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		URL:        URL,
		ParserFunc: parser.ParseCityList,
	})

}

//func determineEncoding(r io.Reader) encoding.Encoding {
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		panic(err)
//	}
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//	return e
//}

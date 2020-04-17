package parser

import (
	"fmt"
	"regexp"

	"github.com/corrots/go-demo/crawler/engine"
)

const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[\w]+)" data-v-[\w]+>([\p{Han}]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	//limit := 0
	reg := regexp.MustCompile(cityListReg)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, val := range matches {
		result.Items = append(result.Items, fmt.Sprintf("City: %s", val[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(val[1]),
			ParserFunc: ParseCity,
		})
		//limit--
		//if limit <= 0 {
		//	break
		//}
	}
	return result
}

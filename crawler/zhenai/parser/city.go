package parser

import (
	"regexp"
	"strings"

	"github.com/corrots/go-demo/crawler/engine"
)

//const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[\w]+)" data-v-[\w]+>([\p{Han}]+)</a>`
const cityReg = `<a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^>]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityReg)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, val := range matches {
		name := string(val[2])
		//result.Items = append(result.Items, "User: "+name)
		url := strings.Replace(string(val[1]), "album", "m", 1)
		result.Requests = append(result.Requests, engine.Request{
			URL: url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}
	return result
}

package parser

import (
	"regexp"
	"strings"

	"github.com/corrots/go-demo/crawler/engine"
)

//const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[\w]+)" data-v-[\w]+>([\p{Han}]+)</a>`
const cityReg = `<a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^>]+)</a>`

func ParseCity(contents []byte, _ string) engine.ParseResult {
	reg := regexp.MustCompile(cityReg)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, val := range matches {
		name := string(val[2])
		//result.Items = append(result.Items, "User: "+name)
		url := strings.Replace(string(val[1]), "album", "m", 1)
		result.Requests = append(result.Requests, engine.Request{
			URL:        url,
			ParserFunc: ProfileParser(name),
		})
	}
	return result
}

func ProfileParser(name string) engine.ParserFunc {
	return func(contents []byte, url string) engine.ParseResult {
		return ParseProfile(contents, name, url)
	}
}

package parser

import (
	"regexp"
	"strconv"

	"github.com/corrots/go-demo/crawler/engine"
	"github.com/corrots/go-demo/crawler/model"
)

var (
	urlReg = regexp.MustCompile(`https://m.zhenai.com/u/(\d+).html`)

	genderReg    = regexp.MustCompile(`<span[^>]+>关注(\p{Han})</span>`)
	marriageReg  = regexp.MustCompile(`<div [^>]*tag[^>]+>([未婚离异丧偶]+)</div>`)
	ageReg       = regexp.MustCompile(`<div [^>]*tag[^>]+>(\d+)岁</div>`)
	xinzuoReg    = regexp.MustCompile(`<div [^>]*tag[^>]+>(\p{Han}+)\([0-9.-]+\)</div>`)
	heightReg    = regexp.MustCompile(`<div [^>]*tag[^>]+>(\d+)cm</div>`)
	weightReg    = regexp.MustCompile(`<div [^>]*tag[^>]+>(\d+)kg</div>`)
	incomeReg    = regexp.MustCompile(`<div [^>]*tag[^>]+>月收入:([^<]+)</div>`)
	educationReg = regexp.MustCompile(`<div [^>]*tag[^>]+>([高中及以下大学本科硕士]+)</div>`)
	hokouReg     = regexp.MustCompile(`<div [^>]*tag[^>]+>籍贯:(\p{Han}+)</div>`)
	carReg       = regexp.MustCompile(`<div [^>]*tag[^>]+>(\p{Han}+车)</div>`)
	houseReg     = regexp.MustCompile(`<div [^>]*tag[^>]+>(\p{Han}+房)</div>`)
)

func ParseProfile(c []byte, name, url string) engine.ParseResult {
	profile := model.Profile{}

	profile.Gender = convertGender(extractString(c, genderReg))
	profile.Name = name
	profile.Age = extractInt(c, ageReg)
	profile.Height = extractInt(c, heightReg)
	profile.Weight = extractInt(c, weightReg)

	profile.Marriage = extractString(c, marriageReg)
	profile.Income = extractString(c, incomeReg)
	profile.Education = extractString(c, educationReg)
	profile.Hokou = extractString(c, hokouReg)
	profile.Xinzuo = extractString(c, xinzuoReg)
	profile.House = extractString(c, houseReg)
	profile.Car = extractString(c, carReg)

	var result engine.ParseResult
	id := extractString([]byte(url), urlReg)
	result.Items = append(result.Items, engine.Item{
		Url:     url,
		Id:      id,
		Payload: profile,
	})
	return result
}

func convertGender(gender string) string {
	if gender == "他" {
		return "男"
	}
	return "女"
}

func extractString(contents []byte, reg *regexp.Regexp) string {
	matches := reg.FindSubmatch(contents)
	if len(matches) >= 2 {
		return string(matches[1])
	}
	return ""
}

func extractInt(contents []byte, reg *regexp.Regexp) int {
	s, err := strconv.Atoi(extractString(contents, reg))
	if err == nil {
		return s
	}
	return 0
}

package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	bytes, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(bytes)
	const resultSize = 470
	expectedURLs := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result Requests should have %d, got %d\n", resultSize, len(result.Requests))
	}

	for i, url := range expectedURLs {
		if result.Requests[i].URL != url {
			t.Errorf("expected URL #%d: %s, got %s\n", i, url, result.Requests[i].URL)
		}
	}
}

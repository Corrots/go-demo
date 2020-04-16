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
	expectedCities := []string{"City: 阿坝", "City: 阿克苏", "City: 阿拉善盟"}
	expectedURLs := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result Requests should have %d, got %d\n", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result Items should have %d, got %d\n", resultSize, len(result.Items))
	}

	for i, url := range expectedURLs {
		if result.Requests[i].URL != url {
			t.Errorf("expected URL #%d: %s, got %s\n", i, url, result.Requests[i].URL)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s, got %s\n", i, city, result.Items[i].(string))
		}
	}
}

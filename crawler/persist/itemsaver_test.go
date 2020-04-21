package persist

import (
	"testing"

	"github.com/corrots/go-demo/crawler/model"
)

func TestSave(t *testing.T) {

	profile := model.Profile{
		Name:      "百羽",
		Gender:    "男",
		Age:       26,
		Height:    185,
		Weight:    95,
		Income:    "3-5千",
		Marriage:  "未婚",
		Education: "高中及以下",
		Hokou:     "四川阿坝",
		Xinzuo:    "魔羯座",
		House:     "打算婚后购房",
		Car:       "未买车",
	}
	err := save(profile)
	if err != nil {
		t.Errorf("err: %v\n", err)
	}

}

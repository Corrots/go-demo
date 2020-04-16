package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMain(m *testing.M) {
	str := `<span data-v-10352ec0="">关注她</span>`

	reg := regexp.MustCompile(`<span[^>]+>关注(\p{Han})</span>`)

	res := reg.FindStringSubmatch(str)
	fmt.Println(res)
}

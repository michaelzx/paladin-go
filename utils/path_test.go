package utils

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {

	test := "asdfasdf"
	fmt.Println(test[0:1])
	var s string
	s = "Go 语言"
	fmt.Println(s[1:4])

	s = "Go 语言"
	rs := []rune(s)
	fmt.Println(string(rs[1:4]))
}

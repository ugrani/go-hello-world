package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	if nil == s{
		return nil
	}
	var stringArray []string = strings.Fields(s)
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}

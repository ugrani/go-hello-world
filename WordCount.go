package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	var stringArray []string = strings.Fields(s)
	var returnMap map[string]int = make (map[string] int)
	for i:=0; i<len(stringArray); i++{
		word := stringArray[i]
		v,ok := returnMap[word]
		if ok{
			returnMap[word] = v+1
		}else{
			returnMap[word] = 1
		}

	}
	return returnMap
}

func main() {
	wc.Test(WordCount)
}

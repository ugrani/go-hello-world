package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {

	var stringArray []string = strings.Fields(s)
	var returnMap map[string]int = make (map[string] int)
	for i:=0; i<len(stringArray); i++{
		word := stringArray[i];
		var count int = len(word)
		fmt.Printf("Count of word %s is %d\n", stringArray[i], count)
		returnMap[word] = count
	}
	return returnMap
}

func main() {
	wc.Test(WordCount)
}

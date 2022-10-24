package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCountMap := make(map[string]int)

	words := strings.Fields(s) // return slice
	count := 0

	for _, word := range words {
		if _, ok := wordCountMap[word]; ok {
			count = wordCountMap[word]
			count++
			wordCountMap[word] = count
		} else {
			wordCountMap[word] = 1
		}
	}

	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}

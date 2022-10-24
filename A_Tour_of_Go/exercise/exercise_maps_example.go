package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	// var m map[string]int
	// makeせずに宣言すると初期値が0ではなくnilになるのでpanic error

	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

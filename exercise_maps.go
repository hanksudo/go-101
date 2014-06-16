package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)

	for _, v := range strings.Fields(s) {
		words[v]++
	}
	return words
}

func main() {
	wc.Test(WordCount)
}

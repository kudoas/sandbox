package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}
	return wordsMap
}

func main() {
	wc.Test(wordCount)
}

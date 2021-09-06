package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := map[string]int{"x": 1}
	//tách chuỗi
	words := strings.Fields(s)
	// lặp các phần tử
	for _, value := range words {
		wordMap[value]++
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}

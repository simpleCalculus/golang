package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var words []string = strings.Fields(s)
	wm := make(map[string]int)	
	for i := range words {
		if elem, ok := wm[words[i]]; ok {
			wm[words[i]] = elem + 1
		} else {
			wm[words[i]] = 1
		}
	}
	return wm
}

func main() {
	wc.Test(WordCount)
}

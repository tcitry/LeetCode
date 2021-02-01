package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}
	var i, j int
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}

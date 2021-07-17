package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("cbacdef"))
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	max_len := 0
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		key := string(s[i])
		if index, ok := m[key]; ok && index >= left {
			left = index + 1
		}
		max_len = max(max_len, i-left+1)
		m[key] = i
	}
	return max_len
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

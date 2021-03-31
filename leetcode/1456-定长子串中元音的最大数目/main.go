package main

import "fmt"

func main() {
	fmt.Println(maxVowels("tryhard", 4))
	fmt.Println(maxVowels("aeiou", 2))
	fmt.Println(maxVowels("leetcode", 3))
}

func maxVowels(s string, k int) int {
	var count, result int
	for i := 0; i < k; i++ {
		if isValid(s[i]) {
			result += 1
		}
	}
	count = result
	for i := 1; i+k <= len(s); i++ {
		if isValid(s[i-1]) {
			count -= 1
		}
		if isValid(s[i+k-1]) {
			count += 1
		}
		if count > result {
			result = count
		}
	}
	return result
}

func isValid(s uint8) bool {
	switch string(s) {
	case "a", "e", "i", "o", "u":
		return true
	default:
		return false
	}
}

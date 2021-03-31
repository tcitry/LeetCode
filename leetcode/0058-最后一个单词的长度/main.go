package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("s"))
}

func lengthOfLastWord(s string) int {
	var i int
	var start, end int
	var finds bool
	for i = len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " && !finds {
			finds = true
			end = i
		}
		if string(s[i]) != " " && finds {
			start = i
		}
		if string(s[i]) == " " && finds {
			break
		}
	}
	if !finds {
		return 0
	}
	return end - start + 1
}

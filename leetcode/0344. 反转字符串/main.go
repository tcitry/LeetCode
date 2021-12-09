package main

import (
	"encoding/json"
	"fmt"
)

// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
func main() {
	v := "hello"
	data, _ := json.Marshal(v)
	reverseString(data)
	fmt.Println(string(data))
}

func reverseString(s []byte) {
	printReverse(0, len(s)-1, s)
}

func printReverse(i, j int, s []byte) {
	if i >= j {
		return
	}
	printReverse(i+1, j-1, s)
	s[i], s[j] = s[j], s[i]
}

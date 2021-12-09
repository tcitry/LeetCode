package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestToChar("hello", []byte("l")[0]))
}

func shortestToChar(S string, C byte) []int {
	var cs, ss, result []int
	var min int

	c := int32(C)
	for index, v := range S {
		if v == c {
			cs = append(cs, index)
		}
		ss = append(ss, index)
	}
	if len(cs) == 0 {
		return []int{}
	}
	for _, i := range ss {
		min = len(ss)
		for _, j := range cs {
			if i < j {
				if j-i < min {
					min = j - i
				}
			} else if i == j {
				min = 0
				break
			} else {
				if i-j < min {
					min = i - j
				}
			}
		}
		result = append(result, min)
	}
	return result
}

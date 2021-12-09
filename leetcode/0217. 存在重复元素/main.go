package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = true
		} else {
			return true
		}
	}
	return false
}

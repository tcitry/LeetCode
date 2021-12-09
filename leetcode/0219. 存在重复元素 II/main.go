package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, n := range nums {
		if v, ok := m[n]; ok {
			if i-v <= k {
				return true
			}
		}
		m[n] = i
	}
	return false
}

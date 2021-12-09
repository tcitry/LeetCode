package main

import "fmt"

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}

// 能用，但感觉不够好
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for index, num := range nums {
		var i int
		if index <= k {
			i = 0
		} else {
			if k > index {
				i = k - index
			} else {
				i = index - k
			}
		}
		for j := index; i < j; i++ {
			if (num-nums[i] <= t && num-nums[i] >= 0) || (num-nums[i] >= -t && num-nums[i] <= 0) {
				if j-i <= k {
					return true
				}
			}
		}
	}
	return false
}

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 5}, 5))
}

func twoSum(nums []int, target int) []int {
	for i, m := range nums {
		for j, n := range nums[i+1:] {
			if m+n == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}

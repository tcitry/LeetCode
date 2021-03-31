package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{1, 2, 3}))
}

func lengthOfLIS(nums []int) int {
	return len(nums)
}

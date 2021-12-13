package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		fmt.Println(dp)
	}

	ret := dp[0]
	for _, v := range dp {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

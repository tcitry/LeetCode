package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(subsets(s))
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	result := [][]int{{}}
	var i int
	j := len(nums)
	for i = 0; i < len(nums); i ++ {
		loopSubSet(nums[0:i], nums[i:j], &result)
	}
	return result
}

func loopSubSet(pre, post []int, result *[][]int) {
	for i := 0; i < len(post); i ++ {
		*result = append(*result, append(pre, i))
	}
}

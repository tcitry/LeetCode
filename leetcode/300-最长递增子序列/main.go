package main

import "fmt"

func main()  {
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))

}

func lengthOfLIS(nums []int) int {
	return len(nums)
}

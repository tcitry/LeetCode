package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3}))
}

func maxProfit(prices []int) int {
	min, max := 0, 0
	if len(prices) <= 1 {
		return 0
	}
	min = prices[0]
	for _, v := range prices {
		if v < min {
			min = v
		}
		if v-min > max {
			max = v - min
		}
	}
	return max
}

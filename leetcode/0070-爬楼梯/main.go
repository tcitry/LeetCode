package main

import "fmt"

func main()  {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	total, pre, post := 0, 0, 1
	for i := 0; i < n ; i ++ {
		total = pre + post
		pre = post
		post = total
	}
	return total
}
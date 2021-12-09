package main

import "fmt"

func main()  {
	fmt.Println(fib(5))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	total, pre, post := 0, 0, 1
	for i := 1; i < n; i ++ {
		total = pre + post
		pre = post
		post = total
	}
	return total
}

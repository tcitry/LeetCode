package main

import "fmt"

func main() {
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(20))
}

func mySqrt(x int) int {
	l := 0
	r := x
	mid := 0
	for l < r {
		mid = (l + r + 1) / 2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

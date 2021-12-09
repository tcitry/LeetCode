package main

import "fmt"

func main() {
	fmt.Println(tribonacci(4))
}

func tribonacci(n int) int {
	if n < 1 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	if n <= 3 {
		return 2
	}
	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}

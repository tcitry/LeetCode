package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, n := range numbers {
		diff := target - n
		if v, in := m[diff]; in {
			return []int{v + 1, i + 1}
		}
		m[n] = i
	}
	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

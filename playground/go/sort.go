package main

import "fmt"

func main() {
	l := []int{10, 2, 3, 23, 11, 20, 1}
	fmt.Println("冒泡:", maopao(l))
}

func maopao(lst []int) []int {
	l := len(lst)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if lst[j] > lst[j+1] {
				lst[j], lst[j+1] = lst[j+1], lst[j]
			}
		}
	}
	return lst
}

func insert(lst []int) []int {
	return lst
}

func selectSort(lst []int) []int {
	return lst
}

func quick(lst []int) []int {
	return lst
}

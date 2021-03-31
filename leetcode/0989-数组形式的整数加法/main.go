package main

import (
	"fmt"
)

func main() {
	fmt.Println(addToArrayForm([]int{0}, 23))
}

func addToArrayForm(A []int, K int) []int {
	var addFlag bool

	var ks []int

	for K/10 > 0 {
		ks = append([]int{K % 10}, ks...)
		K = K / 10
	}
	ks = append([]int{K % 10}, ks...)

	var ll, gl int
	var ls, gs []int
	if len(ks) > len(A) {
		gl = len(ks)
		ll = len(A)
		ls = A
		gs = ks
	} else {
		gl = len(A)
		ll = len(ks)
		ls = ks
		gs = A
	}

	var i, j, item int
	for j = gl - 1; j >= 0; j-- {
		i = j - (gl - ll)
		if i >= 0 {
			item = gs[j] + ls[i]
		} else {
			item = gs[j]
		}
		if addFlag == true {
			item = item + 1
		}
		if item >= 10 {
			gs[j] = item % 10
			addFlag = true
		} else {
			addFlag = false
			gs[j] = item
		}
	}
	if addFlag == true {
		gs = append([]int{1}, gs...)
	}
	return gs
}

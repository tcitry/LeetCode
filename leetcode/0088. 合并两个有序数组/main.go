package main

import "sort"

func main() {
	merge([]int{1, 2, 3}, 1, []int{2, 3, 4}, 2)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

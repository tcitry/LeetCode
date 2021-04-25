package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(arr []int, k int) int {
	// 使用标准库的 container/heap 构造一个最小堆
	h := &IntHeap{}
	heap.Init(h)
	for _, i := range arr {
		heap.Push(h, i)
		fmt.Println(">>>>>>>>>>>", *h)
	}

	var ret int
	for i := 0; i < k; i++ {
		ret = heap.Pop(h).(int)
	}
	return ret
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

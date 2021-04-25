package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
	fmt.Println(getLeastNumbers([]int{1, 0, 1, 2, 1}, 3))

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func getLeastNumbers(arr []int, k int) []int {
	// 使用标准库的 container/heap 构造一个最小堆
	h := &IntHeap{}
	heap.Init(h)
	for _, i := range arr {
		heap.Push(h, i)
	}
	var ret []int
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(h).(int))
	}
	return ret
}

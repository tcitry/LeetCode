package main

import "fmt"

type ListNode struct {
   Val int
   Next *ListNode
}

func main()  {
	fmt.Println(rotateRight(nil, 1))
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n += 1
	}
	rotate := n - k % n
	if rotate == n {
		return head
	}
	iter.Next = head

	for i := 0; i < rotate; i++ {
		iter = iter.Next
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}

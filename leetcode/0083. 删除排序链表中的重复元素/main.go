package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head.Next
	pre := head
	for curr != nil {
		if curr.Val == pre.Val {
			pre.Next = curr.Next
		} else {
			pre = pre.Next
		}
		curr = curr.Next
	}
	return head
}

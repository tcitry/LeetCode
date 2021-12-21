package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	quick := head.Next
	slow := head
	for quick != slow {
		if quick == nil || quick.Next == nil {
			return false
		}
		slow = slow.Next
		quick = quick.Next.Next
	}
	return true
}

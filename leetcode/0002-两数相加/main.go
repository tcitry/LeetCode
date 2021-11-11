package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry bool
	var sum int
	tail := &ListNode{}
	var result *ListNode

	for l1 != nil || l2 != nil {
		if carry {
			sum += 1
			carry = false
		}
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum > 9 {
			sum = sum % 10
			carry = true
		}
		if result == nil {
			result = &ListNode{Val: sum}
			tail = result
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
		sum = 0
	}
	if carry {
		tail.Next = &ListNode{Val: 1}
	}
	return result
}

func main() {
	addTwoNumbers(nil, nil)
}

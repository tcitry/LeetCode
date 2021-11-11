看到题目基本就知道有两种思路，一种是按顺序遍历链表，逐个更新 next 结点，另一种就是利用递归入栈，从最后一个结点开始反向出栈更新 next。

（但是写起来就不是那么顺利了。。）

方法一：迭代

```go
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
```



方法二：递归







## 用哈希表存储

```go
func deleteDuplicates(head *ListNode) *ListNode {
	vm := make(map[int]bool)
	curr := head
	pre := curr
	for curr != nil {
		if _, in := vm[curr.Val]; !in {
			vm[curr.Val] = true
			pre = curr
		} else {
			pre.Next = curr.Next
		}
		curr = curr.Next
	}
	return head
}
```


## 一次遍历

因为是排好序的列表，所以可以直接遍历链表。


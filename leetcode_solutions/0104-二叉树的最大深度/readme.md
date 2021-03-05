### 思路

### 解法 1，深度优先搜索

通过递归，对比结点左右的深度，递归每次得出最大值。

```go
// 深度优先搜索
func maxDepth(root *TreeNode) int {
	if root != nil {
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```


### 解法 2，广度优先搜索



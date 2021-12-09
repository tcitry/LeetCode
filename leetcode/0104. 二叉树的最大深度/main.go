package main

import "fmt"

func main() {
	fmt.Println(maxDepth(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

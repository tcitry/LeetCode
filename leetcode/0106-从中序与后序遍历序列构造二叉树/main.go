package main

func main() {
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	i := 0
	for ; i < len(inorder); i++ {
		if postorder[len(postorder)-1] == inorder[i] {
			break
		}
	}
	root := TreeNode{Val: postorder[len(postorder)-1]}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return &root
}

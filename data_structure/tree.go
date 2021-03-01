package main

import "fmt"

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func main() {
	fmt.Println(">>>>>>>>>>>")
}

func genTree() *Tree {
	return &Tree{}
}

// 先序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
// 中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
// 后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点

// 先序遍历
func preorderRecursion(root *Tree) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderRecursion(root.Left)
	preorderRecursion(root.Right)
}

func preorderTraversal(root *Tree) {
	if root == nil {
		return
	}

}

// 中序遍历
func inorderRecursion(root *Tree) {
	preorderRecursion(root.Left)
	if root == nil {
		return
	}
	preorderRecursion(root.Right)
}

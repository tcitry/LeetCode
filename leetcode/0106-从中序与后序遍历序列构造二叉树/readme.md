举例：
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]


### 解题思路
与根据前序+中序求后序一样，只不过后序的最后一位是根节点，确定后代入递归

### 代码

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
```

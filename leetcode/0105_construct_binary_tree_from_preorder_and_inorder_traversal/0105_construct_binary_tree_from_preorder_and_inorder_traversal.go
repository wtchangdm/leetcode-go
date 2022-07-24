package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

func findIndex(s []int, num int) int {
	for i := range s {
		if s[i] == num {
			return i
		}
	}

	return -1
}

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	// mid acts as the separator of left and right subtree in inorder slice
	// and as the number of elements of the left subtree
	mid := findIndex(inorder, preorder[0])

	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}

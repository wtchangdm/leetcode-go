package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		poped := queue[0]
		queue = queue[1:]

		poped.Left, poped.Right = poped.Right, poped.Left

		if poped.Left != nil {
			queue = append(queue, poped.Left)
		}

		if poped.Right != nil {
			queue = append(queue, poped.Right)
		}
	}

	return root
}

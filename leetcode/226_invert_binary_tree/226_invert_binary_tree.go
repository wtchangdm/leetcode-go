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
		return root
	}

	queue := []*TreeNode{}

	queue = append(queue, root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		current.Left, current.Right = current.Right, current.Left

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return root
}

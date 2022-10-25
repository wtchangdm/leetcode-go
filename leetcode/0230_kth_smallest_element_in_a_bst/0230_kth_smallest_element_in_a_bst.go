package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

func kthSmallestBFS(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}
}

func inOrderDFS(root *TreeNode, nodes []*TreeNode) []*TreeNode {
	if root == nil {
		return nodes
	}

	nodes = inOrderDFS(root.Left, nodes)
	nodes = append(nodes, root)
	nodes = inOrderDFS(root.Right, nodes)

	return nodes
}

func kthSmallestDFS(root *TreeNode, k int) int {
	result := inOrderDFS(root, []*TreeNode{})

	return result[k-1].Val
}

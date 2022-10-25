package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxSubtreeDepth int
	leftSubtreeDepth := maxDepthDFS(root.Left)
	rightSubtreeDepth := maxDepthDFS(root.Right)

	if leftSubtreeDepth > rightSubtreeDepth {
		maxSubtreeDepth = leftSubtreeDepth
	} else {
		maxSubtreeDepth = rightSubtreeDepth
	}

	return 1 + maxSubtreeDepth
}

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		nDrain := len(queue)

		for i := 0; i < nDrain; i++ {
			node := queue[i]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[nDrain:]
		depth++
	}

	return depth
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepthDFSIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	depthStack := []int{1}

	maxDepth := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		depth := depthStack[len(depthStack)-1]
		stack = stack[:len(stack)-1]
		depthStack = depthStack[:len(depthStack)-1]

		if depth > maxDepth {
			maxDepth = depth
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
			depthStack = append(depthStack, depth+1)
		}

		if current.Right != nil {
			stack = append(stack, current.Right)
			depthStack = append(depthStack, depth+1)
		}
	}

	return maxDepth
}

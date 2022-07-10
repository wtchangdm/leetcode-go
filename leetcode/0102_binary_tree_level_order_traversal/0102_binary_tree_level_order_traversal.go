package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		subtrees := []int{}
		nNodes := len(queue)

		for i := 0; i < nNodes; i++ {
			poped := queue[0]
			queue = queue[1:]
			subtrees = append(subtrees, poped.Val)

			if poped.Left != nil {
				queue = append(queue, poped.Left)
			}

			if poped.Right != nil {
				queue = append(queue, poped.Right)
			}
		}

		if len(subtrees) > 0 {
			result = append(result, subtrees)
		}
	}

	return result
}

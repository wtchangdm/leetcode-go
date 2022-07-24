package structs

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BreathFirstSearch(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current.Val)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return result
}

func IsSameTree(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if n1.Val != n2.Val {
		return false
	}

	return IsSameTree(n1.Left, n2.Left) && IsSameTree(n1.Right, n2.Right)
}

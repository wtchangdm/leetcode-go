package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
	current := root

	for current != nil {
		if p.Val > current.Val && q.Val > current.Val {
			current = current.Right
		} else if p.Val < current.Val && q.Val < current.Val {
			current = current.Left
		} else {
			return current
		}
	}

	return nil
}

func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorRecursive(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorRecursive(root.Left, p, q)
	}

	return root
}

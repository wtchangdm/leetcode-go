package leetcode

import (
	"math"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func validateSubBST(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	currentVal := int64(node.Val)
	if currentVal > min && currentVal < max {
		return validateSubBST(node.Left, min, currentVal) && validateSubBST(node.Right, currentVal, max)
	}

	return false
}

func isValidBST(root *TreeNode) bool {
	return validateSubBST(root, math.MinInt64, math.MaxInt64)
}

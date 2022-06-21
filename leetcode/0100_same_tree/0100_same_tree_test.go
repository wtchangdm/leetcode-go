package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestSameTree(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tree2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	if !isSameTree(tree, tree2) {
		t.Errorf("error: %v != %v", tree, tree2)
	}
}

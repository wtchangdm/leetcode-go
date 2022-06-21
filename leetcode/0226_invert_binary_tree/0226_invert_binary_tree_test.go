package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestInvertTree(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	expected := []int{4, 7, 2, 9, 6, 3, 1}
	actual := BreathFirstSearch(invertTree(root))

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("error: %d != %d", expected[i], actual[i])
		}
	}
}

func TestInvertTreeIterative(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	expected := []int{4, 7, 2, 9, 6, 3, 1}
	actual := BreathFirstSearch(invertTreeIterative(root))

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("error: %d != %d", expected[i], actual[i])
		}
	}
}

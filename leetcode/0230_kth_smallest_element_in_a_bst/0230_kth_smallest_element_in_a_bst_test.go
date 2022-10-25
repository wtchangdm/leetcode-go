package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestKthSmallestBFS(t *testing.T) {
	for _, v := range []struct {
		input  *TreeNode
		answer int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			answer: 2,
		},
	} {
		result := kthSmallestBFS(v.input, 2)
		if result != v.answer {
			t.Errorf("error: result %v != %v", result, v.answer)
		}
	}
}

func TestKthSmallestDFS(t *testing.T) {
	for _, v := range []struct {
		input  *TreeNode
		answer int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			answer: 2,
		},
	} {
		result := kthSmallestDFS(v.input, 2)
		if result != v.answer {
			t.Errorf("error: result %v != %v", result, v.answer)
		}
	}
}

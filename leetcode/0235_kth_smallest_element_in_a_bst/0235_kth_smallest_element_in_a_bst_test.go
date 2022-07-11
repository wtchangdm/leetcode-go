package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestKthSmallest(t *testing.T) {
	for _, v := range []struct {
		input  *TreeNode
		target int
		answer int
	}{
		{
			input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			target: 3,
			answer: 3,
		},
	} {
		result := kthSmallest(v.input, v.target)
		if result != v.answer {
			t.Errorf("error: result %v != %v", result, v.answer)
		}
	}
}

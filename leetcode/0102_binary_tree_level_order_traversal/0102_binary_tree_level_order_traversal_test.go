package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestLevelOrder(t *testing.T) {
	for _, v := range []struct {
		input  *TreeNode
		answer [][]int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			answer: [][]int{{3}, {9, 20}, {15, 7}},
		},
	} {
		result := levelOrder(v.input)
		for i := range result {
			for j := range result[i] {
				if result[i][j] != v.answer[i][j] {
					t.Errorf("error: %d != %d", result[i][j], v.answer[i][j])
				}
			}
		}
	}
}

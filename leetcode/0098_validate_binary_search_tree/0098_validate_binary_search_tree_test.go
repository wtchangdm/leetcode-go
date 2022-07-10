package leetcode

import (
	"math"
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestIsValidBST(t *testing.T) {
	for _, v := range []struct {
		input  *TreeNode
		answer bool
	}{
		{
			input: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: -2147483648,
				},
				Right: &TreeNode{
					Val: 2147483647,
				},
			},
			answer: true,
		},
	} {
		result := validateSubBST(v.input, math.MinInt64, math.MaxInt64)
		if result != v.answer {
			t.Errorf("error: result %v != %v", result, v.answer)
		}
	}
}

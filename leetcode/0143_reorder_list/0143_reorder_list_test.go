package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestReorderList(t *testing.T) {
	for _, v := range []struct {
		input  *ListNode
		answer *ListNode
	}{
		{input: BuildListNode([]int{1, 2, 3, 4}), answer: BuildListNode([]int{1, 4, 2, 3})},
		{input: BuildListNode([]int{1, 2, 3, 4, 5}), answer: BuildListNode([]int{1, 5, 2, 4, 3})},
	} {
		reorderList(v.input) // in-place manipulation

		current := v.input
		answer := v.answer
		for current != nil {
			if current.Val != answer.Val {
				t.Errorf("error: %d != %d", current.Val, answer.Val)
			}
			current = current.Next
			answer = answer.Next
		}
	}
}

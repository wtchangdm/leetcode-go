package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestRemoveNthFromEnd(t *testing.T) {
	for _, v := range []struct {
		input  *ListNode
		target int
		answer *ListNode
	}{
		{input: BuildListNode([]int{1, 2, 3, 4, 5}), target: 2, answer: BuildListNode([]int{1, 2, 3, 5})},
		{input: BuildListNode([]int{1}), target: 1, answer: BuildListNode([]int{})},
	} {
		current := removeNthFromEnd(v.input, v.target)

		for current != nil {
			if current.Val != v.answer.Val {
				t.Errorf("error: %d != %d", current.Val, v.answer.Val)
			}
			current = current.Next
			v.answer = v.answer.Next
		}
	}
}

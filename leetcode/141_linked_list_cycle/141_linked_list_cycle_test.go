package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestHasCycle(t *testing.T) {
	node := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node,
	}
	node.Next = node2

	if !hasCycle(node) {
		t.Error("error", node)
	}
}

func TestHasCycleTNH(t *testing.T) {
	node := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node,
	}
	node.Next = node2

	if !hasCycleTNH(node) {
		t.Error("error", node)
	}
}

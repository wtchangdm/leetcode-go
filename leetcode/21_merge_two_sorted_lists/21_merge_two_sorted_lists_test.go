package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	curr := mergeTwoLists(l1, l2)

	for curr != nil && curr.Next != nil {
		if curr.Val > curr.Next.Val {
			t.Errorf("error: %d > %d", curr.Val, curr.Next.Val)
		}
		curr = curr.Next
	}
}

func TestMergeTwoListsRecursive(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	curr := mergeTwoListsRecursive(l1, l2)

	for curr != nil && curr.Next != nil {
		if curr.Val > curr.Next.Val {
			t.Errorf("error: %d > %d", curr.Val, curr.Next.Val)
		}
		curr = curr.Next
	}
}

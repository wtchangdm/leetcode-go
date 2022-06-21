package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestReverseList(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	result := []int{5, 4, 3, 2, 1}

	currentNode := reverseList(node)

	for i := range result {
		if currentNode.Val != result[i] {
			t.Errorf("error: %d != %d", currentNode.Val, result[i])
		}
		currentNode = currentNode.Next
	}
}

func TestReverseListRecursive(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	result := []int{5, 4, 3, 2, 1}

	currentNode := reverseListRecursive(node)

	for i := range result {
		if currentNode.Val != result[i] {
			t.Errorf("error: %d != %d", currentNode.Val, result[i])
		}
		currentNode = currentNode.Next
	}
}
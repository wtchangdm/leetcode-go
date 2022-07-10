package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// we can find the target by shifting two pointers left and right,
	// and keep distance n between them
	dummy := &ListNode{}
	dummy.Next = head
	left, right := dummy, head

	// make right pointer go n steps to make the distance
	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	// at the end of loop, right becomes nil,
	// and left will sit at the previous node of the target
	for right != nil {
		left = left.Next
		right = right.Next
	}

	// update left's Next pointer to skip the target node
	left.Next = left.Next.Next

	// dummy node is an extra node prepended in front of head,
	// and we need to return the head
	return dummy.Next
}

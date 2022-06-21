package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

/**
 * https://leetcode.com/problems/reverse-linked-list/
 * Definition for singly-linked list.
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedSublist := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return reversedSublist
}

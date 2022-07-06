package structs

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(seq []int) *ListNode {
	head := &ListNode{}
	current := head

	for i := range seq {
		current.Val = seq[i]
		if i < len(seq)-1 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return head
}

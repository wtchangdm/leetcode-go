package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

func reorderList(head *ListNode) {
	slow, fast := head, head

	// split list into two half with the last position of the slow pointer
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the right half: "1 -> 2 -> 3 -> 5 -> nil"
	// become:
	// 1 -> 2 -> 3 -> nil and 4 -> 5 -> nil; make "4 -> 5 -> nil" to "nil <- 4 <- 5"
	var prev *ListNode
	current := slow.Next
	slow.Next = nil

	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	// assemble two parts
	current = head
	for prev != nil {
		tmp := current.Next
		current.Next = prev
		current = tmp

		tmp = prev.Next
		prev.Next = current
		prev = tmp
	}
}

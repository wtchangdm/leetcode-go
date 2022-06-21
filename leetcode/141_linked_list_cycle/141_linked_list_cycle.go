package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

// hasCycle uses a map to memoize the visited nodes
// time complexity: O(n)
// space complexity: O(n)
func hasCycle(head *ListNode) bool {
	visited := map[*ListNode]struct{}{}

	curr := head

	for curr != nil {
		if _, found := visited[curr]; found {
			return true
		}

		visited[curr] = struct{}{}
		curr = curr.Next
	}

	return false
}

// hasCycleTNH uses Tortoise and Hare
// time complexity: O(n)
// space complexity: O(1)
func hasCycleTNH(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

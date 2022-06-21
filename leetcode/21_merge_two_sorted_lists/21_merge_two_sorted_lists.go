package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

/**
 * https://leetcode.com/problems/merge-two-sorted-lists/
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	mergedList := &ListNode{}

	l1Current := list1
	l2Current := list2

	mergedHead := mergedList

	for l1Current != nil && l2Current != nil {
		if l1Current.Val < l2Current.Val {
			mergedHead.Next = l1Current
			l1Current = l1Current.Next
		} else {
			mergedHead.Next = l2Current
			l2Current = l2Current.Next
		}
		mergedHead = mergedHead.Next
	}

	if l1Current != nil {
		mergedHead.Next = l1Current
	} else {
		mergedHead.Next = l2Current
	}

	return mergedList.Next
}

func mergeTwoListsRecursive(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursive(list2.Next, list1)
		return list2
	}
}

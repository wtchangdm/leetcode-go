package leetcode

import . "github.com/wtchangdm/leetcode-go/structs"

/**
 * https://leetcode.com/problems/merge-two-sorted-lists/
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	mergedList := &ListNode{} // create a dummy head
	pinnedHead := mergedList  // save merged list head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			mergedList.Next = list1
			list1 = list1.Next
		} else {
			mergedList.Next = list2
			list2 = list2.Next
		}
		mergedList = mergedList.Next
	}

	if list1 != nil {
		mergedList.Next = list1
	} else {
		mergedList.Next = list2
	}

	return pinnedHead.Next // skip the dummy head
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

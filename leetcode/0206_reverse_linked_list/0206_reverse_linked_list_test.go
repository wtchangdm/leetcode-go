package leetcode

import (
	"testing"

	. "github.com/wtchangdm/leetcode-go/structs"
)

func TestReverseList(t *testing.T) {
	node := BuildListNode([]int{1, 2, 3, 4, 5})

	answer := []int{5, 4, 3, 2, 1}

	currentNode := reverseList(node)

	for i := range answer {
		if currentNode.Val != answer[i] {
			t.Errorf("error: %d != %d", currentNode.Val, answer[i])
		}
		currentNode = currentNode.Next
	}
}

func TestReverseListRecursive(t *testing.T) {
	node := BuildListNode([]int{1, 2, 3, 4, 5})

	answer := []int{5, 4, 3, 2, 1}

	currentNode := reverseListRecursive(node)

	for i := range answer {
		if currentNode.Val != answer[i] {
			t.Errorf("error: %d != %d", currentNode.Val, answer[i])
		}
		currentNode = currentNode.Next
	}
}

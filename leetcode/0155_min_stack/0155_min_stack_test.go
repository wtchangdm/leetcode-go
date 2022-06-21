package leetcode

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := &MinStack{}
	minStack.Push(3)
	minStack.Push(2)
	minStack.Push(5)

	if minStack.GetMin() != 2 {
		t.Error("error:", t)
	}
}

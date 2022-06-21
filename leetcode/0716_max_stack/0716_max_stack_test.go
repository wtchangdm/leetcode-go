package leetcode

import (
	"testing"
)

func TestMaxStack(t *testing.T) {
	maxStack := &MaxStack{}
	maxStack.Push(3)
	maxStack.Push(1)
	maxStack.Push(9)

	if maxStack.PeekMax() != 9 {
		t.Errorf("error: %d != %d", maxStack.PeekMax(), 9)
	}
}

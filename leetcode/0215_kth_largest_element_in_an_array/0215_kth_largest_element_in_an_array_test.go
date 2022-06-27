package leetcode

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		kth    int
		answer int
	}{
		{nums: []int{3, 2, 1, 5, 6, 4}, kth: 2, answer: 5},
		{nums: []int{2, 1}, kth: 2, answer: 1},
	} {
		result := findKthLargest(v.nums, v.kth)

		if v.answer != result {
			t.Errorf("error: %d != %d", result, v.answer)
		}
	}
}

package leetcode

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	for _, v := range []struct {
		nums   []int
		kth    int
		result int
	}{
		{nums: []int{3, 2, 1, 5, 6, 4}, kth: 2, result: 5},
		{nums: []int{2, 1}, kth: 2, result: 1},
	} {
		answer := findKthLargest(v.nums, v.kth)

		if v.result != answer {
			t.Errorf("error: %d != %d", answer, v.result)
		}
	}
}

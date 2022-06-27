package leetcode

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	answer := 6

	if maxSubArray(arr) != answer {
		t.Errorf("error: %d != %d", maxSubArray(arr), answer)
	}
}

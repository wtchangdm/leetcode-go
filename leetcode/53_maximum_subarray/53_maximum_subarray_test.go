package leetcode

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := 6

	if maxSubArray(arr) != result {
		t.Errorf("error: %d != %d", maxSubArray(arr), result)
	}
}

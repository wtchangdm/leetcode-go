package leetcode

import "testing"

func TestMissingNumber(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	result := missingNumber(input)
	answer := 5

	if result != answer {
		t.Errorf("error: %d != %d", result, answer)
	}
}

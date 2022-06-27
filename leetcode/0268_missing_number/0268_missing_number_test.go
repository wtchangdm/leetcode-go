package leetcode

import "testing"

func TestMissingNumber(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	answer := missingNumber(input)
	result := 5

	if answer != result {
		t.Errorf("error: %d != %d", answer, result)
	}
}

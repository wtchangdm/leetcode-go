package leetcode

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	for _, v := range []struct {
		input  []int
		answer int
	}{
		{input: []int{3, 4, 5, 1, 2}, answer: 1},
		{input: []int{2, 7, 11, 15}, answer: 2},
		{input: []int{4, 5, 6, 7, 0, 1, 2}, answer: 0},
	} {
		result := findMin(v.input)
		if result != v.answer {
			t.Errorf("error: %d != %d", result, v.answer)
		}
	}
}

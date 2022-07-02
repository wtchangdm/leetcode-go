package leetcode

import (
	"testing"
)

func TestSearch(t *testing.T) {
	for _, v := range []struct {
		input  []int
		target int
		answer int
	}{
		{input: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, answer: 4},
		{input: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, answer: -1},
		{input: []int{1}, target: 0, answer: -1},
	} {
		result := search(v.input, v.target)

		if result != v.answer {
			t.Errorf("error: %d != %d", result, v.answer)
		}
	}
}

package leetcode

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	for _, v := range []struct {
		height []int
		answer int
	}{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, answer: 49},
		{height: []int{1, 2, 1}, answer: 2},
		{height: []int{1, 1}, answer: 1},
	} {
		maxArea := maxArea(v.height)

		if maxArea != v.answer {
			t.Errorf("error: %d != %d", maxArea, v.answer)
		}
	}
}

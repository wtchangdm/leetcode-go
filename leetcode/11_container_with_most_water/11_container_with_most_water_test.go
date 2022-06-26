package leetcode

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	for _, v := range []struct {
		height []int
		result int
	}{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, result: 49},
		{height: []int{1, 2, 1}, result: 2},
		{height: []int{1, 1}, result: 1},
	} {
		maxArea := maxArea(v.height)

		if maxArea != v.result {
			t.Errorf("error: %d != %d", maxArea, v.result)
		}
	}
}

package leetcode

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	for _, v := range []struct {
		input  []int
		answer int
	}{
		{input: []int{7, 1, 5, 3, 6, 4}, answer: 5},
	} {
		if maxProfit(v.input) != v.answer {
			t.Error("error:", v)
		}
	}
}

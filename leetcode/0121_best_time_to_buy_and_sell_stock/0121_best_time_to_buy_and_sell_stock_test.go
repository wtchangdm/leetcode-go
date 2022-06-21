package leetcode

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	for _, v := range []struct {
		input          []int
		expectedResult int
	}{
		{input: []int{7, 1, 5, 3, 6, 4}, expectedResult: 5},
	} {
		if maxProfit(v.input) != v.expectedResult {
			t.Error("error:", v)
		}
	}
}

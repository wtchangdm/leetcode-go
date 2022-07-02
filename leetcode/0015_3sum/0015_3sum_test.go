package leetcode

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	for _, v := range []struct {
		input  []int
		answer [][]int
	}{
		{input: []int{-1, 0, 1, 2, -1, -4}, answer: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{input: []int{}, answer: [][]int{}},
		{input: []int{0}, answer: [][]int{}},
	} {
		result := threeSum(v.input)
		for i := range v.answer {
			for j := range v.answer[i] {
				if v.answer[i][j] != result[i][j] {
					t.Errorf("error: %d != %d", v.answer[i], result[i])
				}
			}
		}
	}
}

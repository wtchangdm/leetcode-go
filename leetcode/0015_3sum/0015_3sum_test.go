package leetcode

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	for _, v := range []struct {
		input          []int
		expectedResult [][]int
	}{
		{input: []int{-1, 0, 1, 2, -1, -4}, expectedResult: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{input: []int{}, expectedResult: [][]int{}},
		{input: []int{0}, expectedResult: [][]int{}},
	} {
		result := threeSum(v.input)
		for i := range v.expectedResult {
			for j := range v.expectedResult[i] {
				if v.expectedResult[i][j] != result[i][j] {
					t.Errorf("error: %d != %d", v.expectedResult[i], result[i])
				}
			}
		}
	}
}

package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	for _, v := range []struct {
		input          []int
		target         int
		expectedResult []int
	}{
		{input: []int{2, 7, 11, 15}, target: 9, expectedResult: []int{0, 1}},
	} {
		result := twoSum(v.input, v.target)
		for i := range v.expectedResult {
			if v.expectedResult[i] != result[i] {
				t.Errorf("error: %d != %d", v.expectedResult[i], result[i])
			}
		}
	}
}

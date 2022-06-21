package leetcode

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	for _, v := range []struct {
		input          []int
		expectedResult bool
	}{
		{input: []int{1, 2, 3, 1}, expectedResult: true},
		{input: []int{1, 2, 3}, expectedResult: false},
	} {

		if containsDuplicate(v.input) != v.expectedResult {
			t.Error("error:", v)
		}
	}
}

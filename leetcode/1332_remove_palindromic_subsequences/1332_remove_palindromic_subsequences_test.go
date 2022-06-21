package leetcode

import "testing"

func TestRemovePalindromeSub(t *testing.T) {
	for _, v := range []struct {
		input          string
		expectedResult int
	}{
		{input: "", expectedResult: 0},
		{input: "a", expectedResult: 1},
		{input: "ababa", expectedResult: 1},
		{input: "abbbaa", expectedResult: 2},
	} {
		if removePalindromeSub(v.input) != v.expectedResult {
			t.Error("error:", v)
		}
	}
}

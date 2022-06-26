package leetcode

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, v := range []struct {
		input  string
		result int
	}{
		{input: "abcabcbb", result: 3},
		{input: "bbbbb", result: 1},
	} {
		longestCount := lengthOfLongestSubstring(v.input)
		if longestCount != v.result {
			t.Errorf("error: %d != %d", longestCount, v.result)
		}
	}
}

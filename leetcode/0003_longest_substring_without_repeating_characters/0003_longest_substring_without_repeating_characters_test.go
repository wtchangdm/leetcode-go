package leetcode

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, v := range []struct {
		input  string
		answer int
	}{
		{input: "abcabcbb", answer: 3},
		{input: "bbbbb", answer: 1},
	} {
		longestCount := lengthOfLongestSubstring(v.input)
		if longestCount != v.answer {
			t.Errorf("error: %d != %d", longestCount, v.answer)
		}
	}
}

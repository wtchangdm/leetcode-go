package leetcode

import "testing"

func TestRemovePalindromeSub(t *testing.T) {
	for _, v := range []struct {
		input  string
		answer int
	}{
		{input: "", answer: 0},
		{input: "a", answer: 1},
		{input: "ababa", answer: 1},
		{input: "abbbaa", answer: 2},
	} {
		if removePalindromeSub(v.input) != v.answer {
			t.Error("error:", v)
		}
	}
}

package leetcode

import (
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	for _, v := range []struct {
		input  string
		answer bool
	}{
		{input: "()[]{}", answer: true},
		{input: "(]", answer: false},
	} {

		if isValidParentheses(v.input) != v.answer {
			t.Error("failed", v)
		}
	}
}

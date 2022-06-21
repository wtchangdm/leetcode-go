package leetcode

import (
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	for _, v := range []struct {
		input  string
		result bool
	}{
		{input: "()[]{}", result: true},
		{input: "(]", result: false},
	} {

		if isValidParentheses(v.input) != v.result {
			t.Error("failed", v)
		}
	}
}

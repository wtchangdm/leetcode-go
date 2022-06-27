package leetcode

import (
	"testing"
)

func TestIsAlphaNumeric(t *testing.T) {
	for _, v := range []struct {
		input  string
		answer bool
	}{
		{input: "A man, a plan, a canal: Panama", answer: true},
		{input: "0P", answer: false},
	} {
		if isPalindrome(v.input) != v.answer {
			t.Error("error:", v)
		}
	}
}

package leetcode

import (
	"testing"
)

func TestIsAlphaNumeric(t *testing.T) {
	for _, v := range []struct {
		input          string
		expectedResult bool
	}{
		{input: "A man, a plan, a canal: Panama", expectedResult: true},
		{input: "0P", expectedResult: false},
	} {
		if isPalindrome(v.input) != v.expectedResult {
			t.Error("error:", v)
		}
	}
}

package leetcode

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	for _, v := range []struct {
		input  string
		answer bool
	}{
		{input: "()[]{}", answer: true},
		{input: "(]", answer: false},
	} {

		if isValid(v.input) != v.answer {
			t.Error("failed", v)
		}
	}
}

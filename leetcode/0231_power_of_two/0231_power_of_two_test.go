package leetcode

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	for _, v := range []struct {
		input  int
		answer bool
	}{
		{input: -5, answer: false},
		{input: 1, answer: true},
		{input: 7, answer: false},
		{input: 8, answer: true},
	} {
		if isPowerOfTwo(v.input) != v.answer {
			t.Errorf("error: answer of %v != %v", v.input, v.answer)
		}
	}
}

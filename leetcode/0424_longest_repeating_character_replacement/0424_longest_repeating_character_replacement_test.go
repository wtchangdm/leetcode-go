package leetcode

import (
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	for _, v := range []struct {
		input  string
		k      int
		answer int
	}{
		{input: "ABAB", k: 2, answer: 4},
		{input: "AABABBA", k: 1, answer: 4},
	} {
		charsToReplace := characterReplacement(v.input, v.k)

		if charsToReplace != v.answer {
			t.Errorf("error: %d != %d", charsToReplace, v.answer)
		}
	}
}

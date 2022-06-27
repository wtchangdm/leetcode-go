package leetcode

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	for _, v := range []struct {
		answer bool
	}{
		{answer: isAnagram("anagram", "nagaram")},
	} {
		if !v.answer {
			t.Error("error:", v.answer)
		}
	}
}

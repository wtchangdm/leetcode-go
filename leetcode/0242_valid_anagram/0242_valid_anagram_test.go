package leetcode

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	for _, v := range []struct {
		result bool
	}{
		{result: isAnagram("anagram", "nagaram")},
	} {
		if !v.result {
			t.Error("error:", v.result)
		}
	}
}

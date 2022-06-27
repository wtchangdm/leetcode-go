package leetcode

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	for _, v := range []struct {
		input  []int
		answer bool
	}{
		{input: []int{1, 2, 3, 1}, answer: true},
		{input: []int{1, 2, 3}, answer: false},
	} {

		if containsDuplicate(v.input) != v.answer {
			t.Error("error:", v)
		}
	}
}

package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	for _, v := range []struct {
		input  []int
		target int
		answer []int
	}{
		{input: []int{2, 7, 11, 15}, target: 9, answer: []int{0, 1}},
	} {
		answer := twoSum(v.input, v.target)
		for i := range v.answer {
			if v.answer[i] != answer[i] {
				t.Errorf("error: %d != %d", v.answer[i], answer[i])
			}
		}
	}
}

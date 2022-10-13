package leetcode

import (
	"sort"
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
		result := twoSum(v.input, v.target)
		sort.Ints(result)
		sort.Ints(v.answer)
		for i := range v.answer {
			if v.answer[i] != result[i] {
				t.Errorf("error: %d != %d", v.answer[i], result[i])
			}
		}
	}
}

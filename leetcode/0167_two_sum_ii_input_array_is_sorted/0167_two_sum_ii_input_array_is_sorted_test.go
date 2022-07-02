package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	for _, v := range []struct {
		input  []int
		target int
		answer []int
	}{
		{input: []int{2, 7, 11, 15}, target: 9, answer: []int{1, 2}},
		{input: []int{2, 3, 4}, target: 6, answer: []int{1, 3}},
	} {
		result := twoSum(v.input, v.target)
		for i := range v.answer {
			if v.answer[i] != result[i] {
				t.Errorf("error: %d != %d", v.answer[i], result[i])
			}
		}
	}
}

func TestTwoSumWithMap(t *testing.T) {
	for _, v := range []struct {
		input  []int
		target int
		answer []int
	}{
		{input: []int{2, 7, 11, 15}, target: 9, answer: []int{1, 2}},
		{input: []int{2, 3, 4}, target: 6, answer: []int{1, 3}},
	} {
		result := twoSumWithMap(v.input, v.target)
		for i := range v.answer {
			if v.answer[i] != result[i] {
				t.Errorf("error: %d != %d", v.answer[i], result[i])
			}
		}
	}
}

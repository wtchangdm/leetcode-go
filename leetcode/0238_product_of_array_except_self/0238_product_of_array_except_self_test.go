package leetcode

import "testing"

func TestProductExceptSelf(t *testing.T) {
	for _, v := range []struct {
		input  []int
		answer []int
	}{
		{input: []int{1, 2, 3, 4}, answer: []int{24, 12, 8, 6}},
	} {
		result := productExceptSelf(v.input)
		for i := range result {
			if result[i] != v.answer[i] {
				t.Errorf("error: %d != %d", result[i], v.answer[i])
			}
		}
	}
}

package leetcode

import "testing"

func TestProductExceptSelf(t *testing.T) {
	for _, v := range []struct {
		input  []int
		result []int
	}{
		{input: []int{1, 2, 3, 4}, result: []int{24, 12, 8, 6}},
	} {
		answer := productExceptSelf(v.input)
		for i := range answer {
			if answer[i] != v.result[i] {
				t.Errorf("error: %d != %d", answer[i], v.result[i])
			}
		}
	}
}

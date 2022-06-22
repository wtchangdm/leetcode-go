package leetcode

import "testing"

func TestLongestConsecutive(t *testing.T) {
	for _, v := range []struct {
		input  []int
		result int
	}{
		{
			input: []int{100, 4, 200, 1, 3, 2}, result: 4,
		},
		{
			input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, result: 9,
		},
	} {
		answer := longestConsecutive(v.input)

		if answer != v.result {
			t.Errorf("error: %d != %d", answer, v.result)
		}
	}
}

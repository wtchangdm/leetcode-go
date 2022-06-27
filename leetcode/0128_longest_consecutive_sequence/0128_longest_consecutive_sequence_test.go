package leetcode

import "testing"

func TestLongestConsecutive(t *testing.T) {
	for _, v := range []struct {
		input  []int
		answer int
	}{
		{
			input: []int{100, 4, 200, 1, 3, 2}, answer: 4,
		},
		{
			input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, answer: 9,
		},
	} {
		result := longestConsecutive(v.input)

		if result != v.answer {
			t.Errorf("error: %d != %d", result, v.answer)
		}
	}
}

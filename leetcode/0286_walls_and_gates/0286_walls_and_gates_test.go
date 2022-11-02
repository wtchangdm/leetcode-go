package leetcode

import (
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	for _, v := range []struct {
		input  [][]int
		answer [][]int
	}{
		{
			input:  [][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}},
			answer: [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}}},
	} {
		wallsAndGates(v.input)
		for i := range v.input {
			for j := range v.input[i] {
				if v.input[i][j] != v.answer[i][j] {
					t.Errorf("%d != %d", v.input[i][j], v.answer[i][j])
				}
			}
		}
	}
}

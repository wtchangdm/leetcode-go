package leetcode

import "testing"

func TestGeneratePascalsTriangle(t *testing.T) {
	input := 5
	answer := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}

	result := generatePascalsTriangle(input)

	for i := range result {
		for j := range result[i] {
			if result[i][j] != answer[i][j] {
				t.Errorf("error: %d != %d", result[i][j], answer[i][j])
			}
		}
	}
}

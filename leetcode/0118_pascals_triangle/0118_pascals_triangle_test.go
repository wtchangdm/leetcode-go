package leetcode

import "testing"

func TestGeneratePascalsTriangle(t *testing.T) {
	input := 5
	expectedResult := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}

	answer := generatePascalsTriangle(input)

	for i := range answer {
		for j := range answer[i] {
			if answer[i][j] != expectedResult[i][j] {
				t.Errorf("error: %d != %d", answer[i][j], expectedResult[i][j])
			}
		}
	}
}

package leetcode

func generatePascalsTriangle(numRows int) [][]int {
	result := [][]int{
		{1},
	}

	if numRows > 1 {
		result = append(result, []int{1, 1})
	}

	for row := 2; row < numRows; row++ {
		result = append(result, make([]int, row+1))
		result[row][0] = 1
		result[row][len(result[row])-1] = 1

		for j := 1; j < len(result[row])-1; j++ {
			result[row][j] = result[row-1][j-1] + result[row-1][j]
		}
	}

	return result
}

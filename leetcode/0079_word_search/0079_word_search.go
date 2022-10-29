package leetcode

func exist(board [][]byte, word string) bool {
	rows, columns := len(board), len(board[0])
	visited := make([][]bool, rows)

	for i := range visited {
		visited[i] = make([]bool, columns)
	}

	var dfs func(row, column, i int) bool
	dfs = func(row, column, i int) bool {
		if i == len(word) {
			return true
		}

		if row < 0 ||
			row >= rows ||
			column < 0 ||
			column >= columns ||
			word[i] != board[row][column] ||
			visited[row][column] {
			return false
		}

		visited[row][column] = true

		result := dfs(row-1, column, i+1) ||
			dfs(row+1, column, i+1) ||
			dfs(row, column-1, i+1) ||
			dfs(row, column+1, i+1)

		visited[row][column] = false

		return result
	}

	for r := range board {
		for c := range board[r] {
			if dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}

package leetcode

func numIslands(grid [][]byte) int {
	islands := 0
	rows, columns := len(grid), len(grid[0])
	visited := make([][]bool, rows)

	for i := range visited {
		visited[i] = make([]bool, columns)
	}

	var bfs = func(r, c int) {
		queue := [][]int{{r, c}}
		visited[r][c] = true

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			directions := [][]int{
				{-1, 0}, // go up
				{1, 0},  // go down
				{0, -1}, // go left
				{0, 1},  // go right
			}

			for _, direction := range directions {
				targetRow := curr[0] + direction[0]
				targetColumn := curr[1] + direction[1]

				if (targetRow >= 0 && targetRow < rows) &&
					(targetColumn >= 0 && targetColumn < columns) &&
					!visited[targetRow][targetColumn] &&
					rune(grid[targetRow][targetColumn]) == '1' {
					visited[targetRow][targetColumn] = true
					queue = append(queue, []int{targetRow, targetColumn})
				}
			}
		}
	}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if !visited[row][column] && rune(grid[row][column]) == '1' {
				bfs(row, column)
				islands += 1
			}
		}
	}

	return islands
}

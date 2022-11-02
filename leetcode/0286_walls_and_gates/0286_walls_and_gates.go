package leetcode

var emptyRoom = (1 << 31) - 1
var gate = 0

var directions = [][]int{
	{1, 0},  // go down
	{-1, 0}, // go up
	{0, 1},  // go right
	{0, -1}, // go left
}

func wallsAndGates(rooms [][]int) {
	rows, cols := len(rooms), len(rooms[0])
	if rows == 0 {
		return
	}

	queue := [][]int{}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if rooms[row][col] == gate {
				queue = append(queue, []int{row, col})
			}
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		row, col := curr[0], curr[1]

		for _, direction := range directions {
			targetRow := row + direction[0]
			targetCol := col + direction[1]

			if targetRow < 0 || targetRow >= rows || targetCol < 0 || targetCol >= cols || rooms[targetRow][targetCol] != emptyRoom {
				continue
			}

			rooms[targetRow][targetCol] = rooms[row][col] + 1
			queue = append(queue, []int{targetRow, targetCol})
		}
	}
}

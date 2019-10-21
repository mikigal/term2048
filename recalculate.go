package main

func recalculate(directory int) bool {
	moved := false
	if directory == Left || directory == Right {
		for row := 0; row < 4; row++ {
			if recalculateRow(directory, row, make([]int, 0), false) {
				moved = true
			}
		}
		return moved
	}

	for col := 0; col < 4; col++ {
		if recalculateCol(directory, col, make([]int, 0), false) {
			moved = true
		}
	}

	return moved
}

func recalculateRow(directory int, row int, modified []int, moved bool) bool {
	if directory == Left {
		for col := 0; col < 3; col++ {
			currentCell := cells[row][col]
			nextCell := cells[row][col+1]

			if col != 3 && currentCell == 0 && cells[row][col+1] != 0 {
				cells[row][col] = cells[row][col+1]
				cells[row][col+1] = 0
				moved = true
				return recalculateRow(directory, row, modified, moved)
			}

			if currentCell != 0 && nextCell != 0 && currentCell == nextCell && !contains(modified, col) && !contains(modified, col+1) {
				currentScore += cells[row][col] * 2
				cells[row][col] *= 2
				cells[row][col+1] = 0
				modified = append(modified, col)
				moved = true
				return recalculateRow(directory, row, modified, moved)
			}
		}
	}

	if directory == Right {
		for col := 3; col > 0; col-- {
			currentCell := cells[row][col]
			nextCell := cells[row][col-1]

			if col != 0 && currentCell == 0 && cells[row][col-1] != 0 {
				cells[row][col] = cells[row][col-1]
				cells[row][col-1] = 0
				moved = true
				return recalculateRow(directory, row, modified, moved)
			}

			if currentCell != 0 && nextCell != 0 && currentCell == nextCell && !contains(modified, col) && !contains(modified, col-1) {
				currentScore += cells[row][col] * 2
				cells[row][col] *= 2
				cells[row][col-1] = 0
				modified = append(modified, col)
				moved = true
				return recalculateRow(directory, row, modified, moved)
			}
		}
	}

	return moved
}

func recalculateCol(directory int, col int, modified []int, moved bool) bool {
	if directory == Up {
		for row := 0; row < 3; row++ {
			currentCell := cells[row][col]
			nextCell := cells[row+1][col]

			if row != 3 && currentCell == 0 && cells[row+1][col] != 0 {
				cells[row][col] = cells[row+1][col]
				cells[row+1][col] = 0
				moved = true
				return recalculateCol(directory, col, modified, moved)
			}

			if currentCell != 0 && nextCell != 0 && currentCell == nextCell && !contains(modified, row) && !contains(modified, row+1) {
				currentScore += cells[row][col] * 2
				cells[row][col] *= 2
				cells[row+1][col] = 0
				modified = append(modified, row)
				moved = true
				return recalculateCol(directory, col, modified, moved)
			}
		}
	}

	if directory == Down {
		for row := 3; row > 0; row-- {
			currentCell := cells[row][col]
			nextCell := cells[row-1][col]

			if row != 0 && currentCell == 0 && cells[row-1][col] != 0 {
				cells[row][col] = cells[row-1][col]
				cells[row-1][col] = 0
				moved = true
				return recalculateCol(directory, col, modified, moved)
			}

			if currentCell != 0 && nextCell != 0 && currentCell == nextCell && !contains(modified, row) && !contains(modified, row-1) {
				currentScore += cells[row][col] * 2
				cells[row][col] *= 2
				cells[row-1][col] = 0
				modified = append(modified, row)
				moved = true
				return recalculateCol(directory, col, modified, moved)
			}
		}
	}

	return moved
}

func isMoveAvailable() bool {
	for row := 0; row < 4; row++ {
		for col := 0; col < 3; col++ {
			currentCell := cells[row][col]
			nextCell := cells[row][col+1]

			if col != 3 && currentCell == 0 && cells[row][col+1] != 0 {
				return true
			}

			if currentCell != 0 && nextCell != 0 && currentCell == nextCell {
				return true
			}
		}

		for col := 3; col > 0; col-- {
			currentCell := cells[row][col]
			nextCell := cells[row][col-1]

			if col != 0 && currentCell == 0 && cells[row][col-1] != 0 {
				return true
			}

			if currentCell != 0 && nextCell != 0 && currentCell == nextCell {
				return true
			}
		}
	}

	for col := 0; col < 4; col++ {
		for row := 0; row < 3; row++ {
			currentCell := cells[row][col]
			nextCell := cells[row+1][col]

			if row != 3 && currentCell == 0 && cells[row+1][col] != 0 {
				return true
			}

			if currentCell != 0 && nextCell != 0 && currentCell == nextCell {
				return true
			}
		}

		for row := 3; row > 0; row-- {
			currentCell := cells[row][col]
			nextCell := cells[row-1][col]

			if row != 0 && currentCell == 0 && cells[row-1][col] != 0 {
				return true
			}

			if currentCell != 0 && nextCell != 0 && currentCell == nextCell {
				return true
			}
		}
	}

	return false
}

func contains(slice []int, col int) bool {
	for _, v := range slice {
		if v == col {
			return true
		}
	}

	return false
}

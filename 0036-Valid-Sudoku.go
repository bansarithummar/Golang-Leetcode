func isValidSudoku(board [][]byte) bool {
    rowSet := [9]map[byte]bool{}
	colSet := [9]map[byte]bool{}
	boxSet := [9]map[byte]bool{}

	for i := 0; i < 9; i++ {
		rowSet[i] = make(map[byte]bool)
		colSet[i] = make(map[byte]bool)
		boxSet[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			val := board[r][c]
			if rowSet[r][val] {
				return false
			}
			rowSet[r][val] = true
			if colSet[c][val] {
				return false
			}
			colSet[c][val] = true
			boxIndex := (r/3)*3 + c/3
			if boxSet[boxIndex][val] {
				return false
			}
			boxSet[boxIndex][val] = true
		}
	}
	return true
}

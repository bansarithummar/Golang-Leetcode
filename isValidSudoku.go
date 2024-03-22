36. Valid Sudoku

func isValidSudoku(board [][]byte) bool {
    rowChecks := make([]map[byte]bool, 9)
    colChecks := make([]map[byte]bool, 9)
    boxChecks := make([]map[byte]bool, 9)
    for i := 0; i < 9; i++ {
        rowChecks[i] = make(map[byte]bool)
        colChecks[i] = make(map[byte]bool)
        boxChecks[i] = make(map[byte]bool)
    }

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            num := board[i][j]
            if num == '.' {
                continue
            }
            boxIndex := (i/3)*3 + j/3

            if rowChecks[i][num] || colChecks[j][num] || boxChecks[boxIndex][num] {
                return false
            }

            rowChecks[i][num] = true
            colChecks[j][num] = true
            boxChecks[boxIndex][num] = true
        }
    }

    return true    
}

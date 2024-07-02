func solveNQueens(n int) [][]string {
    var result [][]string
    board := make([][]byte, n)
    for i := range board {
        board[i] = make([]byte, n)
        for j := range board[i] {
            board[i][j] = '.'
        }
    }
    backtrack(&result, board, 0, n)
    return result
}


func backtrack(result *[][]string, board [][]byte, row, n int) {
    if row == n {
        var solution []string
        for _, line := range board {
            solution = append(solution, string(line))
        }
        *result = append(*result, solution)
        return
    }
    for col := 0; col < n; col++ {
        if isSafe(board, row, col, n) {
            board[row][col] = 'Q'
            backtrack(result, board, row+1, n)
            board[row][col] = '.'
        }
    }
}


func isSafe(board [][]byte, row, col, n int) bool {
    for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    return true   
}

func longestIncreasingPath(matrix [][]int) int {
    ROWS, COLS := len(matrix), len(matrix[0])
    dp := make(map[[2]int]int)

    var dfs func(int, int, int) int
    dfs = func(r, c, prevVal int) int {
        if r < 0 || r >= ROWS || c < 0 || c >= COLS || matrix[r][c] <= prevVal {
            return 0
        }
        if val, found := dp[[2]int{r, c}]; found {
            return val
        }
        res := 1
        res = max(res, 1 + dfs(r+1, c, matrix[r][c]))
        res = max(res, 1 + dfs(r-1, c, matrix[r][c]))
        res = max(res, 1 + dfs(r, c+1, matrix[r][c]))
        res = max(res, 1 + dfs(r, c-1, matrix[r][c]))

        dp[[2]int{r, c}] = res
        return res
    }
    maxPath := 0
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            maxPath = max(maxPath, dfs(r, c, -1))
        }
    }
    return maxPath
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b   
}

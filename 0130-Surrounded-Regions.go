func solve(board [][]byte)  {
    if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
			return
		}
		board[x][y] = '#'
		directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, d := range directions {
			dfs(x+d[0], y+d[1])
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}   
}

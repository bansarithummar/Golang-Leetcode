func updateDistance(grid [][]int) {
	m, n := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			x, y := pos[0]+dir[0], pos[1]+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == math.MaxInt32 {
				grid[x][y] = grid[pos[0]][pos[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}
}

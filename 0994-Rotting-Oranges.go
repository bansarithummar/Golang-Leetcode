func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	queue := [][]int{}
	freshCount := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	minutes := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			pos := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				x, y := pos[0]+dir[0], pos[1]+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					freshCount--
					queue = append(queue, []int{x, y})
				}
			}
		}
		minutes++
	}
	if freshCount > 0 {
		return -1
	}
	return minutes - 1   
}

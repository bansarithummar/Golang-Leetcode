func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range heights {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func([][]int, [][]bool, int, int)
	dfs = func(heights [][]int, ocean [][]bool, i, j int) {
		ocean[i][j] = true
		directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, d := range directions {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !ocean[ni][nj] && heights[ni][nj] >= heights[i][j] {
				dfs(heights, ocean, ni, nj)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(heights, pacific, i, 0)
		dfs(heights, atlantic, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(heights, pacific, 0, j)
		dfs(heights, atlantic, m-1, j)
	}

	result := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result   
}

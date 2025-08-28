func sortMatrix(grid [][]int) [][]int {
    n := len(grid)
	diag := make(map[int][]int) 

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k := i - j
			diag[k] = append(diag[k], grid[i][j])
		}
	}

	for k, a := range diag {
		if k >= 0 {
			sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
		} else {
			sort.Ints(a)
		}
		diag[k] = a
	}

	pos := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k := i - j
			grid[i][j] = diag[k][pos[k]]
			pos[k]++
		}
	}
	return grid    
}

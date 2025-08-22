func minimumArea(grid [][]int) int {
    m := len(grid)
    if m == 0 {
        return 0
    }
    n := len(grid[0])

    minR, minC := m, n
    maxR, maxC := -1, -1

    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == 1 {
                if r < minR { minR = r }
                if r > maxR { maxR = r }
                if c < minC { minC = c }
                if c > maxC { maxC = c }
            }
        }
    }

    if maxR == -1 {
        return 0
    }

    height := maxR - minR + 1
    width  := maxC - minC + 1
    return height * width   
}

func findDiagonalOrder(mat [][]int) []int {
    m, n := len(mat), len(mat[0])
	res := make([]int, 0, m*n)
	for d := 0; d < m+n-1; d++ {
		var tmp []int
		r := 0
		if d >= n {
			r = d - n + 1
		}
		c := d - r
		for r < m && c >= 0 {
			tmp = append(tmp, mat[r][c])
			r++
			c--
		}
		if d%2 == 0 {
			for i := len(tmp) - 1; i >= 0; i-- {
				res = append(res, tmp[i])
			}
		} else {
			res = append(res, tmp...)
		}
	}
	return res  
}

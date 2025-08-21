func numSubmat(mat [][]int) int {
    m := len(mat)
	if m == 0 {
		return 0
	}
	n := len(mat[0])

	heights := make([]int, n)
	var res int64

	type pair struct{ h, cnt int }

	for i := 0; i < m; i++ {
		// Build histogram heights for this row
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}

		// Monotonic stack to count submatrices ending at current row
		stack := make([]pair, 0, n)
		var sum int64 = 0

		for j := 0; j < n; j++ {
			h := heights[j]
			if h == 0 {
				// reset when a zero breaks the histogram
				stack = stack[:0]
				sum = 0
				continue
			}

			cnt := 1
			for len(stack) > 0 && stack[len(stack)-1].h >= h {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				sum -= int64(top.h) * int64(top.cnt)
				cnt += top.cnt
			}
			sum += int64(h) * int64(cnt)
			stack = append(stack, pair{h, cnt})
			res += sum
		}
	}

	return int(res)
    
}

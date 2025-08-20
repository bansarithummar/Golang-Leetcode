func countSquares(matrix [][]int) int {
    m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	// dp[j] will store the size of the largest all-ones square
	// ending at current row i and column j.
	dp := make([]int, n)
	ans := 0

	for i := 0; i < m; i++ {
		prevDiag := 0 // this will hold dp[j-1] from the previous row (top-left)
		for j := 0; j < n; j++ {
			temp := dp[j] // save current dp[j] (which is top for next iteration)
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[j] = 1
				} else {
					// min of (top, left, top-left) + 1
					minNeighbor := dp[j]
					if dp[j-1] < minNeighbor {
						minNeighbor = dp[j-1]
					}
					if prevDiag < minNeighbor {
						minNeighbor = prevDiag
					}
					dp[j] = minNeighbor + 1
				}
				ans += dp[j]
			} else {
				dp[j] = 0
			}
			prevDiag = temp
		}
	}
	return ans
}

func maxValue(events [][]int, k int) int {
    sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	n := len(events)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	// Binary search helper to find latest non-overlapping event
	getPrev := func(i int) int {
		lo, hi := 0, i-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if events[mid][1] < events[i][0] {
				if events[mid+1][1] < events[i][0] {
					lo = mid + 1
				} else {
					return mid + 1
				}
			} else {
				hi = mid - 1
			}
		}
		return 0
	}

	for i := 1; i <= n; i++ {
		prev := getPrev(i - 1)
		for j := 1; j <= k; j++ {
			dp[i][j] = max(dp[i-1][j], dp[prev][j-1]+events[i-1][2])
		}
	}

	return dp[n][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b    
}

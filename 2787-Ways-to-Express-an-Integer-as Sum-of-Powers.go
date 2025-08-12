
func numberOfWays(n int, x int) int {
    const MOD = 1_000_000_007

	pows := make([]int, 0)
	for base := 1; ; base++ {
		val := 1
		for i := 0; i < x; i++ {
			val *= base
			if val > n {
				break
			}
		}
		if val > n {
			if base == 1 { 
				break
			}
			break
		}
		pows = append(pows, val)
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for _, p := range pows {
		for s := n; s >= p; s-- {
			dp[s] += dp[s-p]
			if dp[s] >= MOD {
				dp[s] -= MOD
			}
		}
	}

	return dp[n] 
}

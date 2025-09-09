func peopleAwareOfSecret(n int, delay int, forget int) int {
    const mod = 1_000_000_007
	dp := make([]int, n+1)
	dp[1] = 1
	share := 0

	for day := 2; day <= n; day++ {
		if day-delay >= 1 {
			share = (share + dp[day-delay]) % mod
		}
		if day-forget >= 1 {
			share = (share - dp[day-forget] + mod) % mod
		}
		dp[day] = share
	}

	ans := 0
	for i := n - forget + 1; i <= n; i++ {
		if i >= 1 {
			ans = (ans + dp[i]) % mod
		}
	}
	return ans
}

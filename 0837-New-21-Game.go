func new21Game(n int, k int, maxPts int) float64 {
    if k == 0 || n >= k-1+maxPts {
		return 1.0
	}

	dp := make([]float64, n+1)
	dp[0] = 1.0

	// Wsum keeps sum of last maxPts dp values for states < k
	wsum := 1.0
	ans := 0.0

	for i := 1; i <= n; i++ {
		dp[i] = wsum / float64(maxPts)
		if i < k {
			wsum += dp[i]
		} else {
			// i >= k contributes to final answer (stop drawing)
			ans += dp[i]
		}
		// Slide window: remove dp[i-maxPts] if it was part of wsum
		if i-maxPts >= 0 {
			if i-maxPts < k {
				wsum -= dp[i-maxPts]
			}
		}
	}

	return ans
    
}

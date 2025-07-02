func possibleStringCount(word string, k int) int {
    const MOD = int(1e9 + 7)

	consecutive := []int{1}
	prev := word[0]
	for i := 1; i < len(word); i++ {
		if word[i] == prev {
			consecutive[len(consecutive)-1]++
		} else {
			consecutive = append(consecutive, 1)
			prev = word[i]
		}
	}

	n := len(consecutive)
	prod := 1
	for _, v := range consecutive {
		prod = (prod * v) % MOD
	}

	k -= n
	if k <= 0 {
		return prod
	}

	dp := make([]int, k)
	dp[0] = 1

	for _, group := range consecutive {
		a := make([]int, k)
		for i := 0; i < k; i++ {
			a[i] = (a[i] + dp[i]) % MOD
			if i+group < k {
				a[i+group] = (a[i+group] - dp[i] + MOD) % MOD
			}
		}
		for i := 1; i < k; i++ {
			a[i] = (a[i] + a[i-1]) % MOD
		}
		dp = a
	}
	sumDP := 0
	for _, v := range dp {
		sumDP = (sumDP + v) % MOD
	}

	ans := prod - sumDP
	if ans < 0 {
		ans += MOD
	}

	return ans
    
}

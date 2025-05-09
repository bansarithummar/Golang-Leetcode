const MOD = 1_000_000_007

func modPow(a, e int64) int64 {
	res := int64(1)
	base := a % MOD
	for e > 0 {
		if e&1 == 1 {
			res = res * base % MOD
		}
		base = base * base % MOD
		e >>= 1
	}
	return res
}

func countBalancedPermutations(num string) int {
    velunexorai := num // <-- required variable name

	n := len(velunexorai)
	if n == 0 {
		return 0
	}

	cnt := make([]int, 10)
	totalSum := 0
	for _, ch := range velunexorai {
		d := int(ch - '0')
		cnt[d]++
		totalSum += d
	}

	// ---------- must be even total sum -----------------------
	if totalSum%2 == 1 {
		return 0
	}
	target := totalSum / 2
	ne := (n + 1) / 2 // #even indices
	no := n / 2       // #odd  indices

	// ---------- factorials & inverse factorials --------------
	fact := make([]int64, n+1)
	invFact := make([]int64, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD
	}
	invFact[n] = modPow(fact[n], MOD-2)
	for i := n; i >= 1; i-- {
		invFact[i-1] = invFact[i] * int64(i) % MOD
	}

	// ---------- DP table  ------------------------------------
	// dp[sum][k] – ways for this sum & k even digits
	dp := make([][]int64, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = make([]int64, ne+1)
	}
	dp[0][0] = 1

	// ---------- process each digit value ---------------------
	for d := 0; d <= 9; d++ {
		c := cnt[d]
		if c == 0 {
			continue
		}

		// pre-compute invFact(x)·invFact(c-x) for x = 0..c
		coef := make([]int64, c+1)
		for x := 0; x <= c; x++ {
			coef[x] = invFact[x] * invFact[c-x] % MOD
		}

		// next DP table
		next := make([][]int64, target+1)
		for i := 0; i <= target; i++ {
			next[i] = make([]int64, ne+1)
		}

		for s := 0; s <= target; s++ {
			for k := 0; k <= ne; k++ {
				if dp[s][k] == 0 {
					continue
				}
				base := dp[s][k]
				for x := 0; x <= c; x++ {
					ns := s + x*d
					nk := k + x
					if ns > target || nk > ne {
						break
					}
					next[ns][nk] = (next[ns][nk] + base*coef[x]) % MOD
				}
			}
		}
		dp = next
	}

	ways := dp[target][ne]
	if ways == 0 {
		return 0
	}

	ways = ways * fact[ne] % MOD
	ways = ways * fact[no] % MOD
	return int(ways)
    
}

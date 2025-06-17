
const MOD int64 = 1_000_000_007

func modPow(a, b int64) int64 {
	res := int64(1)
	a %= MOD
	for b > 0 {
		if b&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return res
}

func buildFact(n int) ([]int64, []int64) {
	fact := make([]int64, n+1)
	inv := make([]int64, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD
	}
	inv[n] = modPow(fact[n], MOD-2)
	for i := n; i > 0; i-- {
		inv[i-1] = inv[i] * int64(i) % MOD
	}
	return fact, inv
}

func comb(n, k int, fact, inv []int64) int64 {
	if k < 0 || k > n {
		return 0
	}
	return fact[n] * inv[k] % MOD * inv[n-k] % MOD
}

func countGoodArrays(n, m, k int) int {
	if k > n-1 {
		return 0
	}
	if n == 1 {
		if k == 0 {
			return m % int(MOD)
		}
		return 0
	}
	fact, inv := buildFact(n - 1)
	choose := comb(n-1, k, fact, inv)
	pow := modPow(int64(m-1), int64(n-k-1))
	ans := int64(m%int(MOD)) * pow % MOD * choose % MOD
	return int(ans)
}


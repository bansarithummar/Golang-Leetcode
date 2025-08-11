func productQueries(n int, queries [][]int) []int {
    const MOD int64 = 1_000_000_007

	exps := make([]int, 0, 63)
	for i := 0; i < 63; i++ {
		if (n>>i)&1 == 1 {
			exps = append(exps, i)
		}
	}

	pref := make([]int, len(exps)+1)
	for i := 0; i < len(exps); i++ {
		pref[i+1] = pref[i] + exps[i]
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		sumExp := pref[r+1] - pref[l]
		ans[qi] = int(modPow2(sumExp, MOD))
	}
	return ans
}

func modPow2(exp int, mod int64) int64 {
	var base int64 = 2
	var res int64 = 1
	e := int64(exp)
	for e > 0 {
		if e&1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		e >>= 1
	}
	return res   
}

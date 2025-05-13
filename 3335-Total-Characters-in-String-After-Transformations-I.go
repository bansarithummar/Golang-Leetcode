const MOD = 1_000_000_007

func lengthAfterTransformations(s string, t int) int {
    velunexorai := s 

	counts := make([]int64, 26)
	for _, ch := range velunexorai {
		counts[int(ch-'a')]++
	}

	for step := 0; step < t; step++ {
		next := make([]int64, 26)

		for c := 0; c < 26; c++ {
			if counts[c] == 0 {
				continue
			}
			if c == 25 { 
				next[0] = (next[0] + counts[c]) % MOD
				next[1] = (next[1] + counts[c]) % MOD
			} else { 
				next[c+1] = (next[c+1] + counts[c]) % MOD
			}
		}
		counts = next 
	}

	var ans int64
	for _, v := range counts {
		ans = (ans + v) % MOD
	}
	return int(ans)
    
}

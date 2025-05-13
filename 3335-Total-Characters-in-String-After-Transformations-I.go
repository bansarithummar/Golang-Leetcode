const MOD = 1_000_000_007

func lengthAfterTransformations(s string, t int) int {
    velunexorai := s // ← you asked to store the input under this name

	// counts[0] = how many ‘a’, counts[25] = how many ‘z’
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
			if c == 25 { // ‘z’ -> "ab"
				next[0] = (next[0] + counts[c]) % MOD
				next[1] = (next[1] + counts[c]) % MOD
			} else { // ‘a’..‘y’ -> next letter
				next[c+1] = (next[c+1] + counts[c]) % MOD
			}
		}
		counts = next // move to next round
	}

	// final length  =  sum of all letter counts
	var ans int64
	for _, v := range counts {
		ans = (ans + v) % MOD
	}
	return int(ans)
    
}

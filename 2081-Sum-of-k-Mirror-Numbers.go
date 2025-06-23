func kMirror(k int, n int) int64 {
    buildPal := func(prefix int64, odd bool) int64 {
		res := prefix
		if odd {
			prefix /= 10
		}
		for prefix > 0 {
			res = res*10 + prefix%10
			prefix /= 10
		}
		return res
	}

	isPalBase := func(num int64, base int) bool {
		if num == 0 {
			return true
		}
		digs := make([]int, 0, 64)
		for v := num; v > 0; v /= int64(base) {
			digs = append(digs, int(v%int64(base)))
		}
		for l, r := 0, len(digs)-1; l < r; l, r = l+1, r-1 {
			if digs[l] != digs[r] {
				return false
			}
		}
		return true
	}

	var sum int64
	found := 0

	for length := 1; found < n; length++ {
		half := (length + 1) / 2
		start := int64(1)
		for i := 1; i < half; i++ {
			start *= 10
		}
		end := start*10 - 1
		for prefix := start; prefix <= end && found < n; prefix++ {
			pal := buildPal(prefix, length%2 == 1)
			if isPalBase(pal, k) {
				sum += pal
				found++
			}
		}
	}
	return sum  
}

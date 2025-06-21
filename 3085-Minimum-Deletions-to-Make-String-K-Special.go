func minimumDeletions(word string, k int) int {
    freq := make([]int, 26)
	for _, c := range word {
		freq[c-'a']++
	}
	var f []int
	for _, v := range freq {
		if v > 0 {
			f = append(f, v)
		}
	}
	if len(f) == 0 || k >= len(word) {
		return 0
	}

	sort.Ints(f)                       
	n := len(f)
	pref := make([]int, n+1)           
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + f[i]
	}
	suff := make([]int, n+1)           
	for i := n - 1; i >= 0; i-- {
		suff[i] = suff[i+1] + f[i]
	}

	ans := len(word) 

	for i := 0; i < n; i++ {
		minFreq := f[i]
		target := minFreq + k

		lo, hi := i, n
		for lo < hi {
			mid := (lo + hi) >> 1
			if f[mid] > target {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		r := lo

		delLow := pref[i]                                     
		delHigh := suff[r] - (n-r)*target                     
		if cur := delLow + delHigh; cur < ans {
			ans = cur
		}
	}

	target := k
	lo, hi := 0, n
	for lo < hi {
		mid := (lo + hi) >> 1
		if f[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	r := lo
	delHigh := suff[r] - (n-r)*target
	if delHigh < ans {
		ans = delHigh
	}

	return ans
    
}

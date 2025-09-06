func minOperations(queries [][]int) int64 {
    prefixSumT := func(n int) int64 {
		if n <= 0 {
			return 0
		}
		var total int64
		k := int64(1)
		lo := int64(1)
		hi := int64(4)
		N := int64(n)

		for lo <= N {
			var cnt int64
			if N < hi {
				cnt = N - lo + 1
			} else {
				cnt = hi - lo
			}
			if cnt > 0 {
				total += k * cnt
			}
			k++
			lo = hi
			hi *= 4
		}
		return total
	}

	var ans int64
	for _, q := range queries {
		l, r := q[0], q[1]
		totalHits := prefixSumT(r) - prefixSumT(l-1)
		ans += (totalHits + 1) / 2
	}
	return ans
    
}

func stepsBetween(n, prefix int64) int64 {
	steps, first, next := int64(0), prefix, prefix+1
	for first <= n {
		steps += min(n+1, next) - first 
		first *= 10                     
		next *= 10
	}
	return steps
}

func findKthNumber(n int, k int) int {
    if k == 1 {
		return 1
	}

	cur := int64(1)
	k--            

	n64 := int64(n)
	for k > 0 {
		steps := stepsBetween(n64, cur)
		if int64(k) >= steps {
			k -= int(steps)
			cur++ 
		} else {
			
			k--   
			cur *= 10 
		}
	}
	return int(cur)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b   
}

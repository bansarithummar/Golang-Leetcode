func getNoZeroIntegers(n int) []int {
    isValid := func(x int) bool {
		for x > 0 {
			if x%10 == 0 {
				return false
			}
			x /= 10
		}
		return true
	}

	for a := 1; a < n; a++ {
		b := n - a
		if isValid(a) && isValid(b) {
			return []int{a, b}
		}
	}
	return nil
    
}

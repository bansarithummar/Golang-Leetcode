func makeTheIntegerZero(num1 int, num2 int) int {
    for k := 1; k <= 60; k++ {
		x := int64(num1) - int64(k)*int64(num2)
		if x < int64(k) || x < 0 {
			continue
		}
		if bits.OnesCount64(uint64(x)) <= k {
			return k
		}
	}
	return -1  
}

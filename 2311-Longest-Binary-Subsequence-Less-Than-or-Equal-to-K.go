func abs64(x int64) int64 { if x < 0 { return -x }; return x }

func longestSubsequence(s string, k int) int {
    zeros := 0
	for _, c := range s {
		if c == '0' {
			zeros++
		}
	}

	val, pow := int64(0), int64(1) 
	ones := 0
	for i := len(s) - 1; i >= 0 && pow <= int64(k); i-- {
		if s[i] == '1' && val+pow <= int64(k) {
			val += pow
			ones++
		}
		pow <<= 1
	}

	return zeros + ones  
}

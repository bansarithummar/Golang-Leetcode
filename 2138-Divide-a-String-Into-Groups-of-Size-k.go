func divideString(s string, k int, fill byte) []string {
    rem := len(s) % k
	if rem != 0 {
		s += strings.Repeat(string(fill), k-rem)
	}
	res := make([]string, 0, len(s)/k)
	for i := 0; i < len(s); i += k {
		res = append(res, s[i:i+k])
	}
	return res
    
}

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[rune]bool)
	l := 0
	res := 0

	for r, char := range s {
		for charSet[char] {
			delete(charSet, rune(s[l]))
			l++
		}
		charSet[char] = true
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

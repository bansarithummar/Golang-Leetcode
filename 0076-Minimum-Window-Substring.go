func minWindow(s string, t string) string {
    if len(t) == 0 || len(s) == 0 {
		return ""
	}
	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}
	windowCount := make(map[byte]int)
	required := len(tCount)
	left, right, formed := 0, 0, 0
	ans := [3]int{-1, 0, 0} 

	for right < len(s) {
		c := s[right]
		windowCount[c]++
		
		if tCount[c] > 0 && windowCount[c] == tCount[c] {
			formed++
		}
		for left <= right && formed == required {
			c = s[left]
			if ans[0] == -1 || right-left+1 < ans[0] {
				ans = [3]int{right-left+1, left, right}
			}

			windowCount[c]--
			if tCount[c] > 0 && windowCount[c] < tCount[c] {
				formed--
			}
			left++
		}
		right++
	}

	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]    
}

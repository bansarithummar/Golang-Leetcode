func longestSubarray(nums []int) int {
    l, zeros, best := 0, 0, 0
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			zeros++
		}
		for zeros > 1 {
			if nums[l] == 0 {
				zeros--
			}
			l++
		}
		if r-l > best { 
			best = r - l
		}
	}
	return best 
}

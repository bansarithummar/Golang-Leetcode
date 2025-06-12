func maxAdjacentDistance(nums []int) int {
    n := len(nums)
	if n < 2 { 
		return 0
	}

	maxDiff := 0
	for i := 0; i < n; i++ {
		j := (i + 1) % n                    
		diff := int(math.Abs(float64(nums[i] - nums[j])))
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
    
}

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxLength := 0
	for _, length := range dp {
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b   
}
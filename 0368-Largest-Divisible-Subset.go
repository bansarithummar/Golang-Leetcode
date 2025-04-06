func largestDivisibleSubset(nums []int) []int {
    n := len(nums)
    if n == 0 {
        return []int{}
    }
    
    sort.Ints(nums)
    
    dp := make([]int, n)
    prev := make([]int, n)

    for i := range dp {
        dp[i] = 1
        prev[i] = -1
    }

    maxLen, maxIdx := 1, 0
    
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[i] % nums[j] == 0 && dp[j] + 1 > dp[i] {
                dp[i] = dp[j] + 1
                prev[i] = j
            }
        }
        
        if dp[i] > maxLen {
            maxLen = dp[i]
            maxIdx = i
        }
    }
    
    result := make([]int, 0, maxLen)
    for maxIdx != -1 {
        result = append(result, nums[maxIdx])
        maxIdx = prev[maxIdx]
    }
    
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return result    
}

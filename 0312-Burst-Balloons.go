func maxCoins(nums []int) int {
    n := len(nums)
    
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)
    
    dp := make([][]int, n+2)
    for i := range dp {
        dp[i] = make([]int, n+2)
    }
    
    for length := 1; length <= n; length++ {
        for left := 1; left <= n-length+1; left++ {
            right := left + length - 1
            for i := left; i <= right; i++ {
                dp[left][right] = max(dp[left][right],
                    nums[left-1]*nums[i]*nums[right+1] + dp[left][i-1] + dp[i+1][right])
            }
        }
    }    
    return dp[1][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b   
}

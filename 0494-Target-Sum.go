func findTargetSumWays(nums []int, target int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }
    if (target + totalSum) % 2 != 0 || (target + totalSum) < 0 {
        return 0
    }

    S1 := (target + totalSum) / 2
    dp := make([]int, S1 + 1)
    dp[0] = 1 

    for _, num := range nums {
        for i := S1; i >= num; i-- {
            dp[i] += dp[i - num]
        }
    }
    return dp[S1]    
}

func mostPoints(questions [][]int) int64 {
    n := len(questions)
    dp := make([]int64, n+1)
    
    for i := n - 1; i >= 0; i-- {
        points, jump := questions[i][0], questions[i][1]
        nextIndex := min(i+jump+1, n)
        dp[i] = max(int64(points) + dp[nextIndex], dp[i+1])
    } 
    return dp[0]
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b  
}

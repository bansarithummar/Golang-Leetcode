func soupServings(n int) float64 {
    if n >= 4800 {
        return 1.0 
    }

    dp := make(map[[2]int]float64)

    var dfs func(a, b int) float64
    dfs = func(a, b int) float64 {
        if a <= 0 && b <= 0 {
            return 0.5
        }
        if a <= 0 {
            return 1.0
        }
        if b <= 0 {
            return 0.0
        }

        key := [2]int{a, b}
        if val, ok := dp[key]; ok {
            return val
        }

        res := 0.25 * (
            dfs(a-100, b) +
                dfs(a-75, b-25) +
                dfs(a-50, b-50) +
                dfs(a-25, b-75))

        dp[key] = res
        return res
    }

    return dfs(n, n)  
}

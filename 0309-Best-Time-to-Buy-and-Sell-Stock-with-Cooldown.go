func maxProfit(prices []int) int {
    n := len(prices)
    if n == 0 {
        return 0
    }
    
    hold, sold, cooldown := -prices[0], 0, 0
    
    for i := 1; i < n; i++ {
        newHold := max(hold, cooldown - prices[i])
        newSold := hold + prices[i]
        newCooldown := max(cooldown, sold)
        
        hold = newHold
        sold = newSold
        cooldown = newCooldown
    }
    return max(cooldown, sold)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

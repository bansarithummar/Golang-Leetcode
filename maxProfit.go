121. Best Time to Buy and Sell Stock

func maxProfit(prices []int) int {
	l, r := 0, 1
	maxP := 0
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxP = max(maxP, profit)
		} else {
			l = r
		}
		r++
	}
	return maxP
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


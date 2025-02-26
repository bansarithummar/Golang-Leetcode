func maxAbsoluteSum(nums []int) int {
    maxSum := 0
    minSum := 0
    currentMax := 0
    currentMin := 0
    
    for _, num := range nums {
        currentMax = max(currentMax+num, num)
        maxSum = max(maxSum, currentMax)
        
        currentMin = min(currentMin+num, num)
        minSum = min(minSum, currentMin)
    }
    
    return max(maxSum, -minSum)
}

func max(a, b int) int {
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

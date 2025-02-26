func maxAbsoluteSum(nums []int) int {
    maxSum := 0
    minSum := 0
    currentMax := 0
    currentMin := 0
    
    // One-pass Kadane's algorithm with both max and min tracking
    for _, num := range nums {
        // Update max subarray sum (for positive result)
        currentMax = max(currentMax+num, num)
        maxSum = max(maxSum, currentMax)
        
        // Update min subarray sum (for negative result)
        currentMin = min(currentMin+num, num)
        minSum = min(minSum, currentMin)
    }
    
    // Return the larger absolute value
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

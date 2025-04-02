func maximumTripletValue(nums []int) int64 {
    n := len(nums)
    
    maxDiff := make([]int, n)
    maxSoFar := nums[0]
    
    for j := 1; j < n-1; j++ {
        maxDiff[j] = max(maxDiff[j-1], maxSoFar - nums[j])
        maxSoFar = max(maxSoFar, nums[j])
    }
    
    result := int64(0)
    for k := 2; k < n; k++ {
        value := int64(maxDiff[k-1] * nums[k])
        result = max64(result, value)
    }
    
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func max64(a, b int64) int64 {
    if a > b {
        return a
    }
    return b   
}

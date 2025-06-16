func maximumDifference(nums []int) int {
    minVal := nums[0]
    maxDiff := -1
    for _, v := range nums[1:] {
        if v > minVal {
            if d := v - minVal; d > maxDiff {
                maxDiff = d
            }
        } else if v < minVal {
            minVal = v
        }
    }
    return maxDiff
    
}

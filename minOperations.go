1658. Minimum Operations to Reduce X to Zero


func minOperations(nums []int, x int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    n := len(nums)
    maxLength := -1
    left := 0
    currentSum := 0

    for right := 0; right < n; right++ {
        currentSum += nums[right]
        // Shrink the window from the left if the current sum exceeds the target sum
        for currentSum > total-x && left <= right {
            currentSum -= nums[left]
            left++
        }
        if currentSum == total-x {
            maxLength = max(maxLength, right-left+1)
        }
    } 
    if maxLength == -1 {
        return -1
    }
    return n - maxLength
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

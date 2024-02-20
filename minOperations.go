1658. Minimum Operations to Reduce X to Zero

func minOperations(nums []int, x int) int {
    n := len(nums)
    target := 0
    for _, num := range nums {
        target += num
    }
    target -= x
    
    prefixSum := make(map[int]int)
    prefixSum[0] = -1
    currSum := 0
    maxLength := -1

    for i, num := range nums {
        currSum += num
        if idx, exists := prefixSum[currSum - target]; exists {
            maxLength = max(maxLength, i - idx)
        }
        prefixSum[currSum] = i
    }

    if maxLength == -1 {
        return -1
    } else {
        return n - maxLength
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

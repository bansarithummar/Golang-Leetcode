func maximumUniqueSubarray(nums []int) int {
    seen := make(map[int]bool)
    left, sum, maxSum := 0, 0, 0

    for right := 0; right < len(nums); right++ {
        for seen[nums[right]] {
            seen[nums[left]] = false
            sum -= nums[left]
            left++
        }
        seen[nums[right]] = true
        sum += nums[right]
        if sum > maxSum {
            maxSum = sum
        }
    }

    return maxSum
    
}

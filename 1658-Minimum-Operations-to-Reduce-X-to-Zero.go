func minOperations(nums []int, x int) int {
    totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	target := totalSum - x
	if target < 0 {
		return -1
	}
	n := len(nums)
	currentSum := 0
	left := 0
	maxLength := -1
	for right := 0; right < n; right++ {
		currentSum += nums[right]
		for currentSum > target && left <= right {
			currentSum -= nums[left]
			left++
		}
		if currentSum == target {
			maxLength = max(maxLength, right-left+1)
		}
	}
	if maxLength == -1 {
		return -1
	}
	return n - maxLength   
}

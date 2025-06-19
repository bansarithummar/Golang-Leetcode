func partitionArray(nums []int, k int) int {
    sort.Ints(nums)
	count := 1
	start := nums[0]

	for _, v := range nums {
		if v-start > k {
			count++
			start = v
		}
	}
	return count
    
}

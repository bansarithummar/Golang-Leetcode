func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
	var result [][]int
	var subset []int
	backtrack(nums, &result, subset, 0)
	return result
}

func backtrack(nums []int, result *[][]int, subset []int, start int) {
	temp := make([]int, len(subset))
	copy(temp, subset)
	*result = append(*result, temp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue 
		}
		subset = append(subset, nums[i])
		backtrack(nums, result, subset, i+1)
		subset = subset[:len(subset)-1]
	}    
}

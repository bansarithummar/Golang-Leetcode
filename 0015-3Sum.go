func threeSum(nums []int) [][]int {
    sort.Ints(nums) 
	var result [][]int 

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1 

		for left < right {
			sum := nums[i] + nums[left] + nums[right] 
			switch {
			case sum == 0:
				result = append(result, []int{nums[i], nums[left], nums[right]}) 
				left++
				right--
				
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			case sum < 0:
				left++ 
			default:
				right-- 
			}
		}
	}

	return result
    
}

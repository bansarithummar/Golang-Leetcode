167. Two Sum II - Input Array Is Sorted


func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		curSum := numbers[l] + numbers[r]
		if curSum > target {
			r--
		} else if curSum < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}


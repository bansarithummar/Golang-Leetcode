128. Longest Consecutive Sequence

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	longest := 0
	for _, num := range nums {
		numSet[num] = true
	}
	for _, num := range nums {
		if !numSet[num-1] {
			length := 0
			for numSet[num+length] {
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

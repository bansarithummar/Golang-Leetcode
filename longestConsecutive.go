128. Longest Consecutive Sequence

package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	longest := 0

	// Populate the set with nums
	for _, num := range nums {
		numSet[num] = true
	}

	// Iterate through nums
	for _, num := range nums {
		// If num-1 is not in numSet, start counting consecutive numbers
		if !numSet[num-1] {
			length := 0
			// Count consecutive numbers
			for numSet[num+length] {
				length++
			}
			// Update longest consecutive sequence
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // Output: 4
}


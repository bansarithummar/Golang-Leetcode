2149. Rearrange Array Elements by Sign

package main

import "fmt"

func rearrange(nums []int) []int {
	// Separate positive and negative integers
	var positives []int
	var negatives []int
	for _, num := range nums {
		if num > 0 {
			positives = append(positives, num)
		} else {
			negatives = append(negatives, num)
		}
	}

	// Rearrange the elements following the given conditions
	result := make([]int, 0, len(nums))
	i, j := 0, 0
	for i < len(positives) && j < len(negatives) {
		result = append(result, positives[i], negatives[j])
		i++
		j++
	}

	// Append the remaining elements from the non-empty array
	result = append(result, positives[i:]...)
	result = append(result, negatives[j:]...)

	return result
}

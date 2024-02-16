1481. Least Number of Unique Integers after K Removals

package main

import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	// Count the frequency of each integer
	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
	}
	
	// Store the frequencies in a slice
	frequencies := make([]int, 0, len(freqMap))
	for _, freq := range freqMap {
		frequencies = append(frequencies, freq)
	}
	
	// Sort the frequencies in ascending order
	sort.Ints(frequencies)
	
	// Remove the smallest frequency elements until k elements are removed
	numUnique := len(frequencies)
	for _, freq := range frequencies {
		if k >= freq {
			k -= freq
			numUnique--
		} else {
			break
		}
	}
	
	return numUnique
}

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{5,5,4}, 1))       // Output: 1
	fmt.Println(findLeastNumOfUniqueInts([]int{4,3,1,1,3,3,2}, 3)) // Output: 2
}


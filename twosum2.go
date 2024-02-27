167. Two Sum II - Input Array Is Sorted

package main

import "fmt"

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

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target)) // Output: [1 2]
}


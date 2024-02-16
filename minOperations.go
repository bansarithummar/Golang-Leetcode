1658. Minimum Operations to Reduce X to Zero

package main

import (
	"fmt"
)

func minOperations(nums []int, x int) int {
	target := -x
	for _, num := range nums {
		target += num
	}
	
	if target == 0 {
		return len(nums)
	}
	
	left, right := 0, 0
	sum := 0
	minOps := -1
	
	for right < len(nums) {
		sum += nums[right]
		right++
		
		for sum > target && left < right {
			sum -= nums[left]
			left++
		}
		
		if sum == target {
			if minOps == -1 || len(nums)-right+left < minOps {
				minOps = len(nums) - right + left
			}
		}
	}
	
	if minOps == -1 {
		return -1
	}
	
	return minOps
}

func main() {
	fmt.Println(minOperations([]int{1,1,4,2,3}, 5)) // Output: 2
	fmt.Println(minOperations([]int{5,6,7,8,9}, 4)) // Output: -1
}

11. Container with Most Water

package main

import (
	"fmt"
)

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0

	for l < r {
		area := (r - l) * min(height[l], height[r])
		maxArea = max(area, maxArea)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height)) // Output: 49
}


84. Largest Rectangle in Histogram

package main

import "fmt"

func largestRectangleArea(heights []int) int {
	maxArea := 0
	var stack [][]int

	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			index, height := stack[len(stack)-1][0], stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			maxArea = max(maxArea, height*(i-index))
			start = index
		}
		stack = append(stack, []int{start, h})
	}

	for _, pair := range stack {
		index, height := pair[0], pair[1]
		maxArea = max(maxArea, height*(len(heights)-index))
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

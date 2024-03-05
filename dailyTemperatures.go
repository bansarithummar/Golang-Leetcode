739. Daily Temperatures

package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack [][]int

	for i, t := range temperatures {
		for len(stack) > 0 && t > stack[len(stack)-1][0] {
			stackT, stackInd := stack[len(stack)-1][0], stack[len(stack)-1][1]
			res[stackInd] = i - stackInd
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, []int{t, i})
	}

	return res
}


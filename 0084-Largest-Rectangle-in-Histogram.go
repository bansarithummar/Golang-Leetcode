func largestRectangleArea(heights []int) int {
n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}

	maxArea := 0
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, h*width)
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		h := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		width := n
		if len(stack) > 0 {
			width = n - stack[len(stack)-1] - 1
		}
		maxArea = max(maxArea, h*width)
	}

	return maxArea  
}

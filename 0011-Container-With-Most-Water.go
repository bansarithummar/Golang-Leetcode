func maxArea(height []int) int {
    left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left
		
		currentHeight := height[left]
		if height[right] < height[left] {
			currentHeight = height[right]
		}
		
		currentArea := width * currentHeight
		if currentArea > maxArea {
			maxArea = currentArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

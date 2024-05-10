42. Trapping Rain Water

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	res := 0
	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			res += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			res += rightMax - height[right]
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minEatingSpeed(piles []int, h int) int {
    left, right := 1, max(piles)
	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left   
}

func canFinish(piles []int, k int, h int) bool {
	time := 0
	for _, pile := range piles {
		time += int(math.Ceil(float64(pile) / float64(k)))
	}
	return time <= h
}

func max(piles []int) int {
	maxVal := piles[0]
	for _, pile := range piles {
		if pile > maxVal {
			maxVal = pile
		}
	}
	return maxVal
}

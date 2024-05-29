875. Koko Eating Bananas


func minEatingSpeed(piles []int, h int) int {
	left, right := 1, max(piles)
	res := right

	for left <= right {
		k := (left + right) / 2
		hours := 0
		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(k)))
		}

		if hours <= h {
			res = min(res, k)
			right = k - 1
		} else {
			left = k + 1
		}
	}
	return res
}

func max(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
    n := len(startTime)
	res := 0
	sumDurations := make([]int, n+1)

	for i := 0; i < n; i++ {
		sumDurations[i+1] = sumDurations[i] + endTime[i] - startTime[i]
	}

	for i := k - 1; i < n; i++ {
		right := 0
		if i == n-1 {
			right = eventTime
		} else {
			right = startTime[i+1]
		}

		left := 0
		if i != k-1 {
			left = endTime[i-k]
		}

		freeTime := right - left - (sumDurations[i+1] - sumDurations[i-k+1])
		if freeTime > res {
			res = freeTime
		}
	}

	return res
    
}

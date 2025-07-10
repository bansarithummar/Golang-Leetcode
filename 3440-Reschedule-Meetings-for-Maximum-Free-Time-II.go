func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
    n := len(startTime)
	intervals := make([][2]int, n)
	for i := 0; i < n; i++ {
		intervals[i] = [2]int{startTime[i], endTime[i]}
	}
	
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	gaps := make([]int, 0, n+1)
	gaps = append(gaps, intervals[0][0])
	for i := 1; i < n; i++ {
		gaps = append(gaps, intervals[i][0]-intervals[i-1][1])
	}
	gaps = append(gaps, eventTime-intervals[n-1][1])

	maxBefore := make([]int, n+2)
	maxAfter := make([]int, n+2)
	for i := 1; i <= n; i++ {
		maxBefore[i] = max(maxBefore[i-1], gaps[i-1])
	}
	for i := n; i >= 1; i-- {
		maxAfter[i] = max(maxAfter[i+1], gaps[i])
	}

	maxFree := 0
	for i := 0; i < n; i++ {
		duration := intervals[i][1] - intervals[i][0]
		leftGap := gaps[i]
		rightGap := gaps[i+1]
		combinedGap := leftGap + rightGap
		
		bestGap := 0
		if i > 0 {
			bestGap = max(bestGap, maxBefore[i])
		}
		if i < n-1 {
			bestGap = max(bestGap, maxAfter[i+2])
		}

		if duration <= bestGap {
			maxFree = max(maxFree, combinedGap+duration)
		} else {
			maxFree = max(maxFree, combinedGap)
		}
	}

	for _, gap := range gaps {
		maxFree = max(maxFree, gap)
	}

	return maxFree
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

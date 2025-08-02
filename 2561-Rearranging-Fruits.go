func minCost(basket1 []int, basket2 []int) int64 {
    count := make(map[int]int)
	for _, v := range basket1 {
		count[v]++
	}
	for _, v := range basket2 {
		count[v]++
	}
	// check if possible
	for _, v := range count {
		if v%2 != 0 {
			return -1
		}
	}
	// build surplus lists
	extra1 := []int{}
	extra2 := []int{}
	cnt1 := make(map[int]int)
	cnt2 := make(map[int]int)
	for _, v := range basket1 {
		cnt1[v]++
	}
	for _, v := range basket2 {
		cnt2[v]++
	}
	for k, v := range count {
		diff := cnt1[k] - v/2
		if diff > 0 {
			for i := 0; i < diff; i++ {
				extra1 = append(extra1, k)
			}
		} else if diff < 0 {
			for i := 0; i < -diff; i++ {
				extra2 = append(extra2, k)
			}
		}
	}
	if len(extra1) == 0 {
		return 0
	}
	sort.Ints(extra1)
	sort.Sort(sort.Reverse(sort.IntSlice(extra2)))
	// Find global min
	all := append(basket1, basket2...)
	minValue := all[0]
	for _, v := range all {
		if v < minValue {
			minValue = v
		}
	}
	var cost int64 = 0
	for i := 0; i < len(extra1); i++ {
		a, b := extra1[i], extra2[i]
		if a < b {
			if int64(a) < int64(2*minValue) {
				cost += int64(a)
			} else {
				cost += int64(2 * minValue)
			}
		} else {
			if int64(b) < int64(2*minValue) {
				cost += int64(b)
			} else {
				cost += int64(2 * minValue)
			}
		}
	}
	return cost
    
}

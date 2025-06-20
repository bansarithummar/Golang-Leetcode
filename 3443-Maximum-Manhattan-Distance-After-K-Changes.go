func maxDistance(s string, k int) int {
    abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	ans := 0
	north, south, east, west := 0, 0, 0, 0

	for i, c := range s {
		switch c {
		case 'N':
			north++
		case 'S':
			south++
		case 'E':
			east++
		case 'W':
			west++
		}

		x := abs(north - south)
		y := abs(east - west)
		md := x + y                                 
		extra := min(2*k, (i+1)-md)                 
		if dist := md + extra; dist > ans {
			ans = dist
		}
	}
	return ans
    
}

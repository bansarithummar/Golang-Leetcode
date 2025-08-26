func areaOfMaxDiagonal(dimensions [][]int) int {
    bestD2, bestArea := -1, 0
	for _, d := range dimensions {
		l, w := d[0], d[1]
		d2 := l*l + w*w
		area := l * w
		if d2 > bestD2 || (d2 == bestD2 && area > bestArea) {
			bestD2, bestArea = d2, area
		}
	}
	return bestArea
}

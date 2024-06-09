func carFleet(target int, position []int, speed []int) int {
    n := len(position)
	if n == 0 {
		return 0
	}
	cars := make([][2]float64, n)
	for i := 0; i < n; i++ {
		time := float64(target-position[i]) / float64(speed[i])
		cars[i] = [2]float64{float64(position[i]), time}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})
	fleets := 0
	currentFleetTime := 0.0
	for i := 0; i < n; i++ {
		if cars[i][1] > currentFleetTime {
			fleets++
			currentFleetTime = cars[i][1]
		}
	}
	return fleets    
}

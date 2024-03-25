853. Car Fleet

func carFleet(target int, position []int, speed []int) int {
	type pair struct {
		position int
		speed    int
	}
	var pairs []pair
	for i := 0; i < len(position); i++ {
		pairs = append(pairs, pair{position[i], speed[i]})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].position > pairs[j].position
	})

	var stack []float64
	for _, p := range pairs {
		time := float64(target-p.position) / float64(p.speed)
		stack = append(stack, time)
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}


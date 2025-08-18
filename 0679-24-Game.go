func judgePoint24(cards []int) bool {
    const eps = 1e-6

	nums := make([]float64, 4)
	for i := 0; i < 4; i++ {
		nums[i] = float64(cards[i])
	}

	var dfs func([]float64) bool
	dfs = func(a []float64) bool {
		if len(a) == 1 {
			return abs(a[0]-24.0) < eps
		}
		
		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				x, y := a[i], a[j]

				next := make([]float64, 0, len(a)-1)
				for k := 0; k < len(a); k++ {
					if k != i && k != j {
						next = append(next, a[k])
					}
				}

				candidates := []float64{
					x + y,
					x - y,
					y - x,
					x * y,
				}
				if abs(y) > eps {
					candidates = append(candidates, x/y)
				}
				if abs(x) > eps {
					candidates = append(candidates, y/x)
				}

				for _, v := range candidates {
					next = append(next, v)
					if dfs(next) {
						return true
					}
					next = next[:len(next)-1]
				}
			}
		}
		return false
	}

	return dfs(nums)
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f   
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
    used := make([]bool, len(baskets))
	unplaced := 0

	for _, fruit := range fruits {
		placed := false
		for j := 0; j < len(baskets); j++ {
			if !used[j] && baskets[j] >= fruit {
				used[j] = true
				placed = true
				break
			}
		}
		if !placed {
			unplaced++
		}
	}

	return unplaced
    
}

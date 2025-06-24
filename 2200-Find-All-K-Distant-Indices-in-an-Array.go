func findKDistantIndices(nums []int, key int, k int) []int {
    keyPos := make([]int, 0)
	for i, v := range nums {
		if v == key {
			keyPos = append(keyPos, i)
		}
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	res := make([]int, 0)
	p := 0
	for i := 0; i < len(nums); i++ {
		for p < len(keyPos) && keyPos[p] < i-k {
			p++
		}
		if p < len(keyPos) && abs(keyPos[p]-i) <= k {
			res = append(res, i)
		}
	}
	return res
    
}

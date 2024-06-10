func findLeastNumOfUniqueInts(arr []int, k int) int {
    frequency := make(map[int]int)
	for _, num := range arr {
		frequency[num]++
	}
	frequencies := make([]int, 0, len(frequency))
	for _, freq := range frequency {
		frequencies = append(frequencies, freq)
	}
	sort.Ints(frequencies)
	remainingUniqueInts := len(frequencies)
	for _, freq := range frequencies {
		if k >= freq {
			k -= freq
			remainingUniqueInts--
		} else {
			break
		}
	}
	return remainingUniqueInts   
}

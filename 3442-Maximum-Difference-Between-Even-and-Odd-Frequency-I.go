func maxDifference(s string) int {
    var freq [26]int
	for _, ch := range s {
		freq[ch-'a']++
	}

	maxOdd := -1          
	minEven := 1 << 30    

	for _, f := range freq {
		if f == 0 {
			continue 
		}
		if f&1 == 1 { 
			if f > maxOdd {
				maxOdd = f
			}
		} else { 
			if f < minEven {
				minEven = f
			}
		}
	}

	
	if maxOdd == -1 || minEven == 1<<30 {
		return 0
	}
	return maxOdd - minEven
    
}

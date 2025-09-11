func sortVowels(s string) string {
    isVowel := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}

	// collect vowels
	vowels := []byte{}
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

	// sort vowels by ASCII
	sort.Slice(vowels, func(i, j int) bool { return vowels[i] < vowels[j] })

	// rebuild string
	res := make([]byte, len(s))
	vi := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			res[i] = vowels[vi]
			vi++
		} else {
			res[i] = s[i]
		}
	}

	return string(res)   
}

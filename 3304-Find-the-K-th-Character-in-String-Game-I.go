func kthCharacter(k int) byte {
    word := []byte{'a'}
	for len(word) < k {
		next := make([]byte, len(word))
		for i := 0; i < len(word); i++ {
			if word[i] == 'z' {
				next[i] = 'a'
			} else {
				next[i] = word[i] + 1
			}
		}
		word = append(word, next...)
	}
	return word[k-1]
}

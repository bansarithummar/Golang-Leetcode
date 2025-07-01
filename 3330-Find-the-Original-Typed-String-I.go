func possibleStringCount(word string) int {
    type group struct {
		ch    byte
		count int
	}
	groups := make([]group, 0)
	n := len(word)
	i := 0

	for i < n {
		j := i
		for j < n && word[j] == word[i] {
			j++
		}
		groups = append(groups, group{ch: word[i], count: j - i})
		i = j
	}

	total := 1
	for _, g := range groups {
		if g.count > 1 {
			total += g.count - 1
		}
	}
	return total    
}

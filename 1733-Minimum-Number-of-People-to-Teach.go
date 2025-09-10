func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
    m := len(languages)
	know := make([]map[int]bool, m+1)
	for i := 1; i <= m; i++ {
		know[i] = make(map[int]bool)
		for _, l := range languages[i-1] {
			know[i][l] = true
		}
	}

	// find users in problematic friendships
	problem := make(map[int]bool)
	for _, f := range friendships {
		u, v := f[0], f[1]
		common := false
		for l := range know[u] {
			if know[v][l] {
				common = true
				break
			}
		}
		if !common {
			problem[u], problem[v] = true, true
		}
	}

	if len(problem) == 0 {
		return 0
	}

	// try teaching each language
	res := m
	for lang := 1; lang <= n; lang++ {
		cnt := 0
		for u := range problem {
			if !know[u][lang] {
				cnt++
			}
		}
		if cnt < res {
			res = cnt
		}
	}
	return res  
}

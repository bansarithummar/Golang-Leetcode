func alienOrder(words []string) string {
	graph := make(map[byte][]byte)
	inDegree := make(map[byte]int)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			graph[word[i]] = []byte{}
			inDegree[word[i]] = 0
		}
	}
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]
		minLength := min(len(word1), len(word2))
		for j := 0; j < minLength; j++ {
			if word1[j] != word2[j] {
				graph[word1[j]] = append(graph[word1[j]], word2[j])
				inDegree[word2[j]]++
				break
			}
			if j == minLength-1 && len(word1) > len(word2) {
				return ""
			}
		}
	}
	queue := []byte{}
	for char, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, char)
		}
	}

	result := []byte{}
	for len(queue) > 0 {
		char := queue[0]
		queue = queue[1:]
		result = append(result, char)
		for _, neighbor := range graph[char] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if len(result) != len(inDegree) {
		return ""
	}
	return string(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

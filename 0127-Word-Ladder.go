func differByOne(word1, word2 string) bool {
	diffCount := 0
	for i := range word1 {
		if word1[i] != word2[i] {
			diffCount++
			if diffCount > 1 {
				return false
			}
		}
	}
	return diffCount == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if _, found := wordSet[endWord]; !found {
		return 0 // if endWord is not in the wordList, return 0
	}
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	length := 1

	for len(queue) > 0 {
		nextQueue := []string{}
		for _, word := range queue {
			if word == endWord {
				return length
			}
			for i := range word {
				for c := 'a'; c <= 'z'; c++ {
					newWord := word[:i] + string(c) + word[i+1:]
					if wordSet[newWord] && !visited[newWord] {
						nextQueue = append(nextQueue, newWord)
						visited[newWord] = true
					}
				}
			}
		}
		queue = nextQueue
		length++
	}
	return 0   
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result    
}

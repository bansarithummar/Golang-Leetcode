49.Group Anagrams


func sortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _, str := range strs {
        sortedStr := sortString(str)
        groups[sortedStr] = append(groups[sortedStr], str)
    }

    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }

    return result
    
}

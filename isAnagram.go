242. Valid Anagram

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    countS := make(map[rune]int)
    countT := make(map[rune]int)
    for _, c := range s {
        countS[c]++
    }
    for _, c := range t {
        countT[c]++
    }
    for c, count := range countS {
        if count != countT[c] {
            return false
        }
    }
    return true
}

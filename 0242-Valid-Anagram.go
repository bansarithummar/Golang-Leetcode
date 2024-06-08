func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    counts := make(map[rune]int)
    for _, c := range s {
        counts[c]++
    }
    for _, c := range t {
        counts[c]--
        if counts[c] < 0 {
            return false
        }
    }
    return true
}

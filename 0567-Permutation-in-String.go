func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    s1Count := [26]int{}
    s2Count := [26]int{}
    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]-'a']++
        s2Count[s2[i]-'a']++
    }
    matches := 0
    for i := 0; i < 26; i++ {
        if s1Count[i] == s2Count[i] {
            matches++
        }
    }
    for i := len(s1); i < len(s2); i++ {
        if matches == 26 {
            return true
        }

        right := s2[i] - 'a'
        left := s2[i-len(s1)] - 'a'

        s2Count[right]++
        if s2Count[right] == s1Count[right] {
            matches++
        } else if s2Count[right] == s1Count[right]+1 {
            matches--
        }

        s2Count[left]--
        if s2Count[left] == s1Count[left] {
            matches++
        } else if s2Count[left] == s1Count[left]-1 {
            matches--
        }
    }
    return matches == 26   
}

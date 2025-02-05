func areAlmostEqual(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }

    var diff []int
    for i := range s1 {
        if s1[i] != s2[i] {
            diff = append(diff, i)
        }
        if len(diff) > 2 {
            return false
        }
    }

    return len(diff) == 2 && s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]]
}
    

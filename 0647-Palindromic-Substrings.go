func countSubstrings(s string) int {
    count := 0
    for i := 0; i < len(s); i++ {
        count += expandAroundCenter(s, i, i)
        count += expandAroundCenter(s, i, i+1)
    }
    return count
}

func expandAroundCenter(s string, left int, right int) int {
    count := 0
    for left >= 0 && right < len(s) && s[left] == s[right] {
        count++
        left--
        right++
    }
    return count   
}

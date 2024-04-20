125. Valid Palindrome


func isPalindrome(s string) bool {
    s = strings.ToLower(s) 
    left, right := 0, len(s)-1

    for left < right {
        for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
            left++
        }
        for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
            right--
        }
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

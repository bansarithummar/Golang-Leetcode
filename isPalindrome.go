125. Valid Palindrome

func isPalindrome(s string) bool {
    s = strings.ToLower(s) // Convert to lowercase
    left, right := 0, len(s)-1

    for left < right {
        // Move left index forward if not alphanumeric
        for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
            left++
        }
        // Move right index backward if not alphanumeric
        for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
            right--
        }
        // Compare characters
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

125. Valid Palindrome

package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlphaNum(rune(s[r])) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s)) // Output: true
}


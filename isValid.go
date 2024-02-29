20. Valid Parentheses

package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)
	closeToOpen := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, c := range s {
		if _, ok := closeToOpen[c]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == closeToOpen[c] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}


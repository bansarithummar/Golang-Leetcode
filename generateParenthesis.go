22. Generate Parentheses

package main

import "fmt"

func generateParenthesis(n int) []string {
	var res []string
	var stack []rune

	var backtrack func(openN, closedN int)
	backtrack = func(openN, closedN int) {
		if openN == n && closedN == n {
			res = append(res, string(stack))
			return
		}

		if openN < n {
			stack = append(stack, '(')
			backtrack(openN+1, closedN)
			stack = stack[:len(stack)-1]
		}

		if closedN < openN {
			stack = append(stack, ')')
			backtrack(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return res
}

func main() {
	n := 3
	result := generateParenthesis(n)
	fmt.Println("Result:", result)
}

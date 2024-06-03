func generateParenthesis(n int) []string {
    var result []string
	var current string
	backtrack(n, 0, 0, current, &result)
	return result
}

func backtrack(n, open, close int, current string, result *[]string) {
	if len(current) == 2*n {
		*result = append(*result, current)
		return
	}
	if open < n {
		backtrack(n, open+1, close, current+"(", result)
	}
	if close < open {
		backtrack(n, open, close+1, current+")", result)
	}    
}

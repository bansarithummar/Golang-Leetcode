22. Generate Parentheses


func generateParenthesis(n int) []string {
var result []string
    backtrack(&result, "", 0, 0, n)
    return result
}

func backtrack(result *[]string, cur string, open, close, max int) {
    if len(cur) == max*2 {
        *result = append(*result, cur)
        return
    }
    
    if open < max {
        backtrack(result, cur+"(", open+1, close, max)
    }
    if close < open {
        backtrack(result, cur+")", open, close+1, max)
    }   
}

func clearDigits(s string) string {
    if len(s) == 0 {
        return ""
    }
    
    chars := []rune(s)
    stack := []rune{}
    
    for _, c := range chars {
        if !isDigit(c) {
            stack = append(stack, c)
        } else if len(stack) > 0 {
            // Remove closest non-digit to left
            stack = stack[:len(stack)-1]
        }
    }
    
    return string(stack)
}

func isDigit(r rune) bool {
    return r >= '0' && r <= '9'   
}

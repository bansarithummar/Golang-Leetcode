func reverse(x int) int {
    const maxInt = 2147483647    // 2^31 - 1
    const minInt = -2147483648   // -2^31

    // Result to store reversed number
    result := 0

    // Keep going until x becomes 0
    for x != 0 {
        // Get last digit
        digit := x % 10
        if result > maxInt/10 || result < minInt/10 {
            return 0
        }
        
        // Multiply current result by 10 and add digit
        result = result * 10 + digit
        
        // Remove last digit from x
        x = x / 10
    }
    
    return result
    
}

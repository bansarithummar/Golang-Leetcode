func reverse(x int) int {
    const maxInt = 2147483647    
    const minInt = -2147483648   

    result := 0
    for x != 0 {
        digit := x % 10
        if result > maxInt/10 || result < minInt/10 {
            return 0
        }
        result = result * 10 + digit
        x = x / 10
    }
    return result    
}

func isHappy(n int) bool {
    seen := make(map[int]bool)

    sumOfSquares := func(num int) int {
        sum := 0
        for num > 0 {
            digit := num % 10
            sum += digit * digit
            num /= 10
        }
        return sum
    }

    for n != 1 && !seen[n] {
        seen[n] = true
        n = sumOfSquares(n)
    }

    return n == 1
    
}

func getSum(a int, b int) int {
    for b != 0 {
        // Calculate the carry
        carry := (a & b) << 1
        // Calculate the sum without carry
        a = a ^ b
        // Update b to be the carry
        b = carry
    }
    return a
    
}

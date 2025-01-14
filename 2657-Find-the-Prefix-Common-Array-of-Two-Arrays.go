func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    result := make([]int, n)
    seen := make(map[int]bool) 
    common := 0               // Tracks the count of common elements

    for i := 0; i < n; i++ {
        if seen[A[i]] {
            common++
        } else {
            seen[A[i]] = true
        }
        if seen[B[i]] {
            common++
        } else {
            seen[B[i]] = true
        }
        result[i] = common
    }

    return result
    
}

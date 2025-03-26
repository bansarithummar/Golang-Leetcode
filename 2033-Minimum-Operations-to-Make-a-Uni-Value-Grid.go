func minOperations(grid [][]int, x int) int {
    // Flatten the grid into a single slice
    values := flatten(grid)
    
    // Sort the values
    sort.Ints(values)
    
    // Check transformability
    if !isTransformable(values, x) {
        return -1
    }
    
    // Find median
    median := values[len(values)/2]
    
    // Calculate total operations
    return calculateOperations(values, median, x)
}

// Flatten 2D grid to 1D slice
func flatten(grid [][]int) []int {
    var values []int
    for _, row := range grid {
        values = append(values, row...)
    }
    return values
}

// Check if all values can be transformed
func isTransformable(values []int, x int) bool {
    base := values[0]
    for _, val := range values {
        if abs(val - base) % x != 0 {
            return false
        }
    }
    return true
}

// Calculate minimum operations to transform to median
func calculateOperations(values []int, median int, x int) int {
    totalOps := 0
    for _, val := range values {
        totalOps += abs(val - median) / x
    }
    return totalOps
}

// Absolute value helper function
func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

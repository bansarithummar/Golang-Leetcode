func minOperations(grid [][]int, x int) int {
    values := flatten(grid)
    sort.Ints(values)
    if !isTransformable(values, x) {
        return -1
    }
    median := values[len(values)/2]
    return calculateOperations(values, median, x)
}

func flatten(grid [][]int) []int {
    var values []int
    for _, row := range grid {
        values = append(values, row...)
    }
    return values
}

func isTransformable(values []int, x int) bool {
    base := values[0]
    for _, val := range values {
        if abs(val - base) % x != 0 {
            return false
        }
    }
    return true
}

func calculateOperations(values []int, median int, x int) int {
    totalOps := 0
    for _, val := range values {
        totalOps += abs(val - median) / x
    }
    return totalOps
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

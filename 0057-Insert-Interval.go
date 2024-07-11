func insert(intervals [][]int, newInterval []int) [][]int {
    var result [][]int
    inserted := false    
    for _, interval := range intervals {
        if interval[1] < newInterval[0] {
            result = append(result, interval)
        } else if interval[0] > newInterval[1] {
            if !inserted {
                result = append(result, newInterval)
                inserted = true
            }
            result = append(result, interval)
        } else {
            newInterval[0] = min(newInterval[0], interval[0])
            newInterval[1] = max(newInterval[1], interval[1])
        }
    }
    if !inserted {
        result = append(result, newInterval)
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b    
}

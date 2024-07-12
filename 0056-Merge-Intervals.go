func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    result := [][]int{intervals[0]}
    for _, interval := range intervals[1:] {
        last := result[len(result)-1]
        if interval[0] <= last[1] {
            last[1] = max(last[1], interval[1])
        } else {
            result = append(result, interval)
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b    
}

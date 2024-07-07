func leastInterval(tasks []byte, n int) int {
    taskCount := make(map[byte]int)
    maxFreq := 0
    for _, task := range tasks {
        taskCount[task]++
        if taskCount[task] > maxFreq {
            maxFreq = taskCount[task]
        }
    }
    maxCount := 0
    for _, count := range taskCount {
        if count == maxFreq {
            maxCount++
        }
    }
    partCount := maxFreq - 1
    partLength := n - (maxCount - 1)
    emptySlots := partCount * partLength
    availableTasks := len(tasks) - maxFreq*maxCount
    idles := max(0, emptySlots - availableTasks)

    return len(tasks) + idles
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b    
}

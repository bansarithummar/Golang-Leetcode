func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
    slices.Sort(tasks)
    slices.Sort(workers)

    return sort.Search(len(tasks)+1, func(n int) bool {
        if len(workers) < n {
            return true
        }

        return !canAssignTasks(tasks[:n], workers[len(workers)-n:], pills, strength)
    }) - 1
}

func canAssignTasks(tasks []int, workers []int, pills int, strength int) bool {
    availableTasks := []int{} 
    taskIdx := 0            

    for workerIdx := 0; workerIdx < len(workers); workerIdx++ {
        workerStrength := workers[workerIdx]

        for taskIdx < len(tasks) && workerStrength+strength >= tasks[taskIdx] {
            availableTasks = append(availableTasks, tasks[taskIdx])
            taskIdx++
        }

        if len(availableTasks) == 0 {
            return false
        }

        if workerStrength >= availableTasks[0] {
            availableTasks = availableTasks[1:]
            continue
        }

        if pills > 0 {
            pills--
            availableTasks = availableTasks[:len(availableTasks)-1] // Assign hardest feasible task
            continue
        }

        return false
    }

    return true   
}

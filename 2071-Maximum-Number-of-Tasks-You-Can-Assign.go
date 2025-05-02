func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
    slices.Sort(tasks)
    slices.Sort(workers)

    // Binary search to find the maximum number of assignable tasks
    return sort.Search(len(tasks)+1, func(n int) bool {
        if len(workers) < n {
            return true
        }
        // Check if it's possible to assign the hardest 'n' tasks
        return !canAssignTasks(tasks[:n], workers[len(workers)-n:], pills, strength)
    }) - 1
}

func canAssignTasks(tasks []int, workers []int, pills int, strength int) bool {
    availableTasks := []int{} // Queue of tasks a worker can attempt
    taskIdx := 0              // Index of the current task to be checked

    for workerIdx := 0; workerIdx < len(workers); workerIdx++ {
        workerStrength := workers[workerIdx]

        // Add all tasks that can potentially be done by this worker (with a pill)
        for taskIdx < len(tasks) && workerStrength+strength >= tasks[taskIdx] {
            availableTasks = append(availableTasks, tasks[taskIdx])
            taskIdx++
        }

        // No suitable task available
        if len(availableTasks) == 0 {
            return false
        }

        // If worker can complete the easiest task without a pill
        if workerStrength >= availableTasks[0] {
            availableTasks = availableTasks[1:]
            continue
        }

        // Worker must use a pill to handle the hardest remaining feasible task
        if pills > 0 {
            pills--
            availableTasks = availableTasks[:len(availableTasks)-1] // Assign hardest feasible task
            continue
        }

        // Unable to assign any task to this worker
        return false
    }

    return true
    
}

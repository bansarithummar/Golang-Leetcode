func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    inDegree := make([]int, numCourses)
    
    for _, pair := range prerequisites {
        course, prereq := pair[0], pair[1]
        graph[prereq] = append(graph[prereq], course)
        inDegree[course]++
    }   
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }    
    count := 0
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        count++
        for _, next := range graph[current] {
            inDegree[next]--
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    return count == numCourses   
}

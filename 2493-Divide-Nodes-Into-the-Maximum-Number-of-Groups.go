func magnificentSets(n int, edges [][]int) int {
    graph := make(map[int][]int)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }

    color := make([]int, n+1) 
    components := [][]int{}

    var bfs func(int) ([]int, bool)
    bfs = func(start int) ([]int, bool) {
        queue := []int{start}
        color[start] = 1
        component := []int{start}

        for len(queue) > 0 {
            node := queue[0]
            queue = queue[1:]

            for _, neighbor := range graph[node] {
                if color[neighbor] == 0 {
                    color[neighbor] = -color[node]
                    queue = append(queue, neighbor)
                    component = append(component, neighbor)
                } else if color[neighbor] == color[node] {
                    return nil, false 
                }
            }
        }
        return component, true
    }

    for i := 1; i <= n; i++ {
        if color[i] == 0 {
            component, valid := bfs(i)
            if !valid {
                return -1
            }
            components = append(components, component)
        }
    }

    totalGroups := 0
    for _, component := range components {
        totalGroups += bfsMaxDepth(graph, component)
    }

    return totalGroups
}

func bfsMaxDepth(graph map[int][]int, component []int) int {
    maxDepth := 0
    for _, node := range component {
        depth := bfsDepth(graph, node)
        if depth > maxDepth {
            maxDepth = depth
        }
    }
    return maxDepth
}

func bfsDepth(graph map[int][]int, start int) int {
    queue := []int{start}
    visited := make(map[int]bool)
    visited[start] = true
    depth := 0

    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            for _, neighbor := range graph[node] {
                if !visited[neighbor] {
                    visited[neighbor] = true
                    queue = append(queue, neighbor)
                }
            }
        }
        depth++
    }
    return depth
    
}

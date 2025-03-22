func countCompleteComponents(n int, edges [][]int) int {
    graph := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }
    
    visited := make([]bool, n)
    completeCount := 0
    
    for i := 0; i < n; i++ {
        if !visited[i] {
            vertexCount, edgeCount := 0, 0
            
            dfs(i, graph, visited, &vertexCount, &edgeCount)
            if edgeCount == vertexCount * (vertexCount - 1) {
                completeCount++
            }
        }
    }
    
    return completeCount
}

func dfs(node int, graph [][]int, visited []bool, vertexCount *int, edgeCount *int) {
    visited[node] = true
    *vertexCount++
    *edgeCount += len(graph[node])
    
    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            dfs(neighbor, graph, visited, vertexCount, edgeCount)
        }
    }
    
}

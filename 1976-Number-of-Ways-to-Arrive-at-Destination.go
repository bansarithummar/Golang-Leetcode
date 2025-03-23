func countPaths(n int, roads [][]int) int {
    graph := make([][][2]int, n)
    for _, road := range roads {
        u, v, time := road[0], road[1], road[2]
        graph[u] = append(graph[u], [2]int{v, time})
        graph[v] = append(graph[v], [2]int{u, time})
    }
    
    // Initialize distance and ways arrays
    const mod = 1_000_000_007
    dist := make([]int64, n)
    ways := make([]int, n)
    
    // Set initial values
    for i := 0; i < n; i++ {
        dist[i] = math.MaxInt64
    }
    dist[0] = 0
    ways[0] = 1
    
    // Priority queue for Dijkstra's algorithm
    pq := &PriorityQueue{}
    heap.Push(pq, &Item{node: 0, distance: 0})
    
    // Dijkstra's algorithm
    for pq.Len() > 0 {
        item := heap.Pop(pq).(*Item)
        node, distance := item.node, item.distance
        
        // Skip if we've found a better path already
        if int64(distance) > dist[node] {
            continue
        }
        
        // Process neighbors
        for _, neighbor := range graph[node] {
            nextNode, time := neighbor[0], neighbor[1]
            newDist := dist[node] + int64(time)
            
            if newDist < dist[nextNode] {
                // Found a shorter path
                dist[nextNode] = newDist
                ways[nextNode] = ways[node]
                heap.Push(pq, &Item{node: nextNode, distance: int(newDist)})
            } else if newDist == dist[nextNode] {
                // Found another path with the same length
                ways[nextNode] = (ways[nextNode] + ways[node]) % mod
            }
        }
    }
    
    return ways[n-1]
}

type Item struct {
   node     int
   distance int
   index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int) {
   pq[i], pq[j] = pq[j], pq[i]
   pq[i].index = i
   pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
   n := len(*pq)
   item := x.(*Item)
   item.index = n
   *pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
   old := *pq
   n := len(old)
   item := old[n-1]
   old[n-1] = nil
   item.index = -1
   *pq = old[0 : n-1]
   return item
}

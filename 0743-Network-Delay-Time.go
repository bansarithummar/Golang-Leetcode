type Edge struct {
	target, weight int
}

type MinHeap []Node

type Node struct {
	vertex, dist int
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func networkDelayTime(times [][]int, n int, k int) int {
    graph := make(map[int][]Edge)
	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		graph[u] = append(graph[u], Edge{v, w})
	}
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	minHeap := &MinHeap{Node{k, 0}}
	heap.Init(minHeap)
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(Node)
		u, d := node.vertex, node.dist
		if d > dist[u] {
			continue
		}
		for _, edge := range graph[u] {
			v, w := edge.target, edge.weight
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				heap.Push(minHeap, Node{v, dist[v]})
			}
		}
	}
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1 
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	return maxDist   
}

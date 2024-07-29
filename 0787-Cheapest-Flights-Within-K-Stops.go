type Flight struct {
	city, cost, stops int
}

type MinHeap []Flight

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Flight))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make(map[int][][]int)
	for _, flight := range flights {
		from, to, price := flight[0], flight[1], flight[2]
		graph[from] = append(graph[from], []int{to, price})
	}
	pq := &MinHeap{}
	heap.Push(pq, Flight{city: src, cost: 0, stops: 0})
	costs := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		costs[i] = make(map[int]int)
	}
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Flight)

		if curr.city == dst {
			return curr.cost
		}
		if curr.stops > k {
			continue
		}
		for _, neighbor := range graph[curr.city] {
			nextCity, price := neighbor[0], neighbor[1]
			nextCost := curr.cost + price

			if _, ok := costs[nextCity][curr.stops+1]; !ok || nextCost < costs[nextCity][curr.stops+1] {
				costs[nextCity][curr.stops+1] = nextCost
				heap.Push(pq, Flight{city: nextCity, cost: nextCost, stops: curr.stops + 1})
			}
		}
	}
	return -1
}

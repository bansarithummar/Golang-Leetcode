type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func maxEvents(events [][]int) int {
    
    sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	i, res, day := 0, 0, 0
	n := len(events)

	for i < n || h.Len() > 0 {
		if h.Len() == 0 {
			day = events[i][0]
		}
		for i < n && events[i][0] == day {
			heap.Push(h, events[i][1])
			i++
		}
		for h.Len() > 0 && (*h)[0] < day {
			heap.Pop(h)
		}
		if h.Len() > 0 {
			heap.Pop(h)
			res++
			day++
		}
	}

	return res
}

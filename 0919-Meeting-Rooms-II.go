type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func minMeetingDays(intervals []Interval) int {
    if len(intervals) == 0 {
        return 0
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    minHeap := &MinHeap{}
    heap.Init(minHeap)

    for _, interval := range intervals {
        if minHeap.Len() > 0 && (*minHeap)[0] <= interval.Start {
            heap.Pop(minHeap)
        }
        heap.Push(minHeap, interval.End)
    }
    return minHeap.Len()
}

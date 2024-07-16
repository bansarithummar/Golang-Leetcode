type Interval struct {
    Start int
    End   int
}

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return (h[i][1] - h[i][0]) < (h[j][1] - h[j][0]) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func minInterval(intervals [][]int, queries []int) []int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    queryWithIndex := make([][2]int, len(queries))
    for i, q := range queries {
        queryWithIndex[i] = [2]int{q, i}
    }
    sort.Slice(queryWithIndex, func(i, j int) bool {
        return queryWithIndex[i][0] < queryWithIndex[j][0]
    })
    result := make([]int, len(queries))
    for i := range result {
        result[i] = -1 
    }
    minHeap := &MinHeap{}
    heap.Init(minHeap)
    i := 0
    for _, qw := range queryWithIndex {
        query, idx := qw[0], qw[1]
        for i < len(intervals) && intervals[i][0] <= query {
            heap.Push(minHeap, intervals[i])
            i++
        }
        for minHeap.Len() > 0 && (*minHeap)[0][1] < query {
            heap.Pop(minHeap)
        }
        if minHeap.Len() > 0 {
            smallest := (*minHeap)[0]
            result[idx] = smallest[1] - smallest[0] + 1
        }
    }
    return result    
}

347. Top K Frequent Elements

type pair struct {
	num  int
	freq int
}

type maxHeap []pair

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	h := &maxHeap{}
	heap.Init(h)
	for num, freq := range counter {
		heap.Push(h, pair{num, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(pair).num
	}
	return result
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type MaxHeap struct {
    IntHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }

type MedianFinder struct {
    maxHeap *MaxHeap
    minHeap *IntHeap    
}

func Constructor() MedianFinder {
    maxH := &MaxHeap{IntHeap: IntHeap{}}
    minH := &IntHeap{}
    heap.Init(maxH)
    heap.Init(minH)
    return MedianFinder{
        maxHeap: maxH,
        minHeap: minH,
    }    
}

func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.maxHeap, num)
    heap.Push(this.minHeap, heap.Pop(this.maxHeap))

    if this.maxHeap.Len() < this.minHeap.Len() {
        heap.Push(this.maxHeap, heap.Pop(this.minHeap))
    }    
}

func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len() > this.minHeap.Len() {
        return float64(this.maxHeap.IntHeap[0])
    }
    return float64(this.maxHeap.IntHeap[0]+(*this.minHeap)[0]) / 2.0    
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

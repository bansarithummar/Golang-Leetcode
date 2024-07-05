type Point struct {
    x, y int
    dist float64
}

type PointHeap []Point

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].dist > h[j].dist }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
    *h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func kClosest(points [][]int, k int) [][]int {
    h := &PointHeap{}
    heap.Init(h)
    
    for _, point := range points {
        x, y := point[0], point[1]
        dist := math.Sqrt(float64(x*x + y*y))
        heap.Push(h, Point{x, y, dist})
        if h.Len() > k {
            heap.Pop(h)
        }
    } 
    result := make([][]int, k)
    for i := 0; i < k; i++ {
        p := heap.Pop(h).(Point)
        result[i] = []int{p.x, p.y}
    }
    return result   
}

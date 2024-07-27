type Cell struct {
	row, col, time int
}

type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}


func swimInWater(grid [][]int) int {
    n := len(grid)
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	minHeap := &MinHeap{Cell{0, 0, grid[0][0]}}
	heap.Init(minHeap)
	visited[0][0] = true
	for minHeap.Len() > 0 {
		cell := heap.Pop(minHeap).(Cell)
		r, c, t := cell.row, cell.col, cell.time
		if r == n-1 && c == n-1 {
			return t
		}
		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] {
				visited[nr][nc] = true
				heap.Push(minHeap, Cell{nr, nc, max(t, grid[nr][nc])})
			}
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

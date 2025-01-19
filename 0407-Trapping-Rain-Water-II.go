type Cell struct {
	row, col, height int
}

type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
    if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}

	m, n := len(heightMap), len(heightMap[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	h := &MinHeap{}
	heap.Init(h)

	// Add boundary cells to the heap
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(h, Cell{i, j, heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	waterTrapped := 0

	for h.Len() > 0 {
		cell := heap.Pop(h).(Cell)
		for _, dir := range directions {
			nr, nc := cell.row+dir[0], cell.col+dir[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] {
				visited[nr][nc] = true
				waterTrapped += int(math.Max(0, float64(cell.height-heightMap[nr][nc])))
				heap.Push(h, Cell{nr, nc, int(math.Max(float64(cell.height), float64(heightMap[nr][nc])))})
			}
		}
	}
	return waterTrapped
}

func manhattanDistance(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type UnionFind struct {
	parent, rank []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) 
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
		return true
	}
	return false
}

func minCostConnectPoints(points [][]int) int {
    n := len(points)
	edges := make([][3]int, 0, n*(n-1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := manhattanDistance(points[i], points[j])
			edges = append(edges, [3]int{dist, i, j})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})
	uf := NewUnionFind(n)
	minCost := 0
	for _, edge := range edges {
		if uf.Union(edge[1], edge[2]) {
			minCost += edge[0]
		}
	}
	return minCost    
}

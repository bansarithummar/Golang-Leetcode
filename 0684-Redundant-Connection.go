type UnionFind struct {
    parent []int
    rank   []int
}

func NewUnionFind(size int) *UnionFind {
    parent := make([]int, size)
    rank := make([]int, size)
    for i := 0; i < size; i++ {
        parent[i] = i
        rank[i] = 1
    }
    return &UnionFind{parent: parent, rank: rank}
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

    if rootX == rootY {
        return false 
    }

    // Union by rank
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

func findRedundantConnection(edges [][]int) []int {
    uf := NewUnionFind(len(edges) + 1) // +1 because nodes are 1-indexed
    
    for _, edge := range edges {
        if !uf.Union(edge[0]-1, edge[1]-1) { // Subtract 1 to make the nodes 0-indexed
            return edge
        }
    }
    
    return nil 
}

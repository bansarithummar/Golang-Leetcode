type UnionFind struct {
    parent []int
    rank   []int
}

func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    rank := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
        rank[i] = 1
    }
    return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Find(node int) int {
    if uf.parent[node] != node {
        uf.parent[node] = uf.Find(uf.parent[node])
    }
    return uf.parent[node]
}

func (uf *UnionFind) Union(node1, node2 int) bool {
    root1 := uf.Find(node1)
    root2 := uf.Find(node2)
    
    if root1 == root2 {
        return false
    }
    
    if uf.rank[root1] > uf.rank[root2] {
        uf.parent[root2] = root1
    } else if uf.rank[root1] < uf.rank[root2] {
        uf.parent[root1] = root2
    } else {
        uf.parent[root2] = root1
        uf.rank[root1]++
    }
    
    return true
}

func validTree(n int, edges [][]int) bool {
    // A valid tree must have exactly n-1 edges
    if len(edges) != n-1 {
        return false
    }
    uf := NewUnionFind(n)
    for _, edge := range edges {
        if !uf.Union(edge[0], edge[1]) {
            return false // Found a cycle
        }
    }    
    return true
}

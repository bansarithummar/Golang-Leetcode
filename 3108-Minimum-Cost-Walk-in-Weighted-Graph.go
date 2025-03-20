func findRoot(parent []int, node int) int {
	if parent[node] != node {
		parent[node] = findRoot(parent, parent[node]) // Path compression
	}
	return parent[node]
}

func minimumCost(n int, edges [][]int, query [][]int) []int {
    parent := make([]int, n)
	minPathCost := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		minPathCost[i] = -1
	}

	for _, edge := range edges {
		source, target, weight := edge[0], edge[1], edge[2]
		sourceRoot := findRoot(parent, source)
		targetRoot := findRoot(parent, target)

		minPathCost[targetRoot] &= weight

		if sourceRoot != targetRoot {
			minPathCost[targetRoot] &= minPathCost[sourceRoot]
			parent[sourceRoot] = targetRoot
		}
	}

	result := make([]int, len(query))
	for i, q := range query {
		start, end := q[0], q[1]
		if start == end {
			result[i] = 0
		} else if findRoot(parent, start) != findRoot(parent, end) {
			result[i] = -1
		} else {
			result[i] = minPathCost[findRoot(parent, start)]
		}
	}

	return result
    
}

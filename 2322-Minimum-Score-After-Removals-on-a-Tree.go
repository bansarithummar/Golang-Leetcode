func minimumScore(nums []int, edges [][]int) int {
    reachTime, leaveTime, subtreeXors := dfs(nums, edges)

	result := math.MaxInt
	for i := 1; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			xor1 := subtreeXors[0]
			xor2 := subtreeXors[i]
			xor3 := subtreeXors[j]

			if inTree(i, j, reachTime, leaveTime) {
				xor1 ^= subtreeXors[i]
				xor2 ^= subtreeXors[j]
			} else if inTree(j, i, reachTime, leaveTime) {
				xor1 ^= subtreeXors[j]
				xor3 ^= subtreeXors[i]
			} else {
				xor1 ^= subtreeXors[i] ^ subtreeXors[j]
			}
			result = min(result, max(xor1, xor2, xor3)-min(xor1, xor2, xor3))
		}
	}
	return result
}

func dfs(nums []int, edges [][]int) (reachTime, leaveTime, subtreeXors []int) {
	adjacencyList := buildAdjacencyList(edges)
	time := 0
	reachTime = make([]int, len(nums))
	leaveTime = make([]int, len(nums))
	subtreeXors = make([]int, len(nums))

	var dfsHelper func(root int, parent int)
	dfsHelper = func(root int, parent int) {
		reachTime[root] = time
		time++
		subtreeXors[root] = nums[root]

		for _, child := range adjacencyList[root] {
			if child == parent {
				continue
			}
			dfsHelper(child, root)
			subtreeXors[root] ^= subtreeXors[child]
		}

		leaveTime[root] = time
		time++
	}

	dfsHelper(0, -1)
	return
}

func buildAdjacencyList(edges [][]int) [][]int {
	result := make([][]int, len(edges)+1)
	for _, edge := range edges {
		result[edge[0]] = append(result[edge[0]], edge[1])
		result[edge[1]] = append(result[edge[1]], edge[0])
	}
	return result
}

func inTree(
	root int,
	node int,
	reachTime []int,
	leaveTime []int,
) bool {
	return reachTime[root] < reachTime[node] && leaveTime[node] < leaveTime[root]
}

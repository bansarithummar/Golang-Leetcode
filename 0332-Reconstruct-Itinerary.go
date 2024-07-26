func findItinerary(tickets [][]string) []string {
    graph := make(map[string][]string)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		graph[from] = append(graph[from], to)
	}

	for from := range graph {
		sort.Strings(graph[from])
	}
	var result []string
	var dfs func(string)

	dfs = func(airport string) {
		for len(graph[airport]) > 0 {
			next := graph[airport][0]
			graph[airport] = graph[airport][1:]
			dfs(next)
		}
		result = append(result, airport)
	}
	dfs("JFK")

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result   
}

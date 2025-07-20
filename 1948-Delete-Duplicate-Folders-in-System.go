type Node struct {
	children map[string]*Node
	name     string
	isDup    bool
}

func deleteDuplicateFolder(paths [][]string) [][]string {
    root := &Node{children: make(map[string]*Node)}

	// Step 1: Build folder tree
	for _, path := range paths {
		curr := root
		for _, folder := range path {
			if curr.children[folder] == nil {
				curr.children[folder] = &Node{
					children: make(map[string]*Node),
					name:     folder,
				}
			}
			curr = curr.children[folder]
		}
	}

	// Step 2: Serialize subtrees and mark duplicates
	serialMap := map[string][]*Node{}
	var serialize func(*Node) string
	serialize = func(node *Node) string {
		if len(node.children) == 0 {
			return ""
		}
		parts := []string{}
		for name, child := range node.children {
			parts = append(parts, name+"("+serialize(child)+")")
		}
		sort.Strings(parts)
		serial := strings.Join(parts, "")
		serialMap[serial] = append(serialMap[serial], node)
		return serial
	}
	serialize(root)

	for _, nodes := range serialMap {
		if len(nodes) > 1 {
			for _, node := range nodes {
				node.isDup = true
			}
		}
	}

	var res [][]string
	var dfs func(*Node, []string)
	dfs = func(node *Node, path []string) {
		if node != root {
			path = append(path, node.name)
			if !node.isDup {
				res = append(res, append([]string{}, path...))
			}
		}
		for _, child := range node.children {
			if !child.isDup {
				dfs(child, path)
			}
		}
	}
	dfs(root, []string{})

	return res
    
}

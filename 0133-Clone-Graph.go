/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    visited := make(map[*Node]*Node)
    var clone func(*Node) *Node
    clone = func(n *Node) *Node {
        if n == nil {
            return nil
        }
        if _, ok := visited[n]; ok {
            return visited[n]
        }
        copy := &Node{Val: n.Val}
        visited[n] = copy
        for _, neighbor := range n.Neighbors {
            copy.Neighbors = append(copy.Neighbors, clone(neighbor))
        }       
        return copy
    }    
    return clone(node)
}

func printGraph(node *Node) {
    visited := make(map[*Node]bool)
    var dfs func(*Node)
    dfs = func(n *Node) {
        if n == nil || visited[n] {
            return
        }
        visited[n] = true
        fmt.Printf("Node %d: ", n.Val)
        for _, neighbor := range n.Neighbors {
            fmt.Printf("%d ", neighbor.Val)
        }
        fmt.Println()
        for _, neighbor := range n.Neighbors {
            dfs(neighbor)
        }
    }
    dfs(node)   
}

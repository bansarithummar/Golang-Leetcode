138. Copy List with Random Pointer


func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    nodeMap := make(map[*Node]*Node)
    curr := head
    for curr != nil {
        nodeMap[curr] = &Node{Val: curr.Val}
        curr = curr.Next
    }
    curr = head
    for curr != nil {
        newNode := nodeMap[curr]
        newNode.Next = nodeMap[curr.Next]
        newNode.Random = nodeMap[curr.Random]
        curr = curr.Next
    }
    return nodeMap[head]    
}

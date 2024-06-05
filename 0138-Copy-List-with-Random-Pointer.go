/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
		return nil
	}
	current := head
	for current != nil {
		newNode := &Node{Val: current.Val, Next: current.Next}
		current.Next = newNode
		current = newNode.Next
	}
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}
	current = head
	copyHead := head.Next
	for current != nil {
		copyNode := current.Next
		current.Next = copyNode.Next
		if copyNode.Next != nil {
			copyNode.Next = copyNode.Next.Next
		}
		current = current.Next
	}
	return copyHead
}

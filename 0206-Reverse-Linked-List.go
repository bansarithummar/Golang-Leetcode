206. Reverse Linked List


func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
	current := head

	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}

	return prev
}

19. Remove Nth Node From End of List


func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Val: 0, Next: head}
    first := head
    length := 0
    for first != nil {
        length++
        first = first.Next
    }
    length -= n
    first = dummy
    for length > 0 {
        length--
        first = first.Next
    }
    first.Next = first.Next.Next
    return dummy.Next
}

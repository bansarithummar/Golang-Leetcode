2. Add Two Numbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    current := dummyHead
    carry := 0

    for l1 != nil || l2 != nil {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        carry = sum / 10
        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
    }
    if carry > 0 {
        current.Next = &ListNode{Val: carry}
    }
    return dummyHead.Next
}

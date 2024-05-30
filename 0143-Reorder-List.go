143. Reorder List


func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    middle := slow.Next
    slow.Next = nil
    var prev *ListNode
    for middle != nil {
        next := middle.Next
        middle.Next = prev
        prev = middle
        middle = next
    }
    reversed := prev
    current := head
    for current != nil && reversed != nil {
        next1 := current.Next
        next2 := reversed.Next
        current.Next = reversed
        reversed.Next = next1
        current = next1
        reversed = next2
    }    
}

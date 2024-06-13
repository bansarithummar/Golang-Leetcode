/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k == 1 {
        return head
    }
    reverse := func(start *ListNode, end *ListNode) *ListNode {
        var prev *ListNode
        curr := start
        next := start.Next
        for curr != end {
            curr.Next, prev, curr, next = prev, curr, next, next.Next
        }
        return prev
    }
    count := 0
    temp := head
    for temp != nil {
        count++
        temp = temp.Next
    }
    dummy := &ListNode{0, head}
    prevGroupEnd := dummy

    for count >= k {
        curr := prevGroupEnd.Next
        nextGroupStart := curr.Next
        newHead := reverse(curr, nextGroupStart)
        prevGroupEnd.Next = newHead
        curr.Next = nextGroupStart
        prevGroupEnd = curr
        count -= k
    }
    return dummy.Next
}

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
		for curr != end {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		return prev
	}
	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy

	for {
		groupStart := prevGroupEnd.Next
		groupEnd := prevGroupEnd
		for i := 0; i < k && groupEnd != nil; i++ {
			groupEnd = groupEnd.Next
		}
		if groupEnd == nil {
			break
		}
		nextGroupStart := groupEnd.Next
		reversedGroupStart := reverse(groupStart, nextGroupStart)
		prevGroupEnd.Next = reversedGroupStart
		groupStart.Next = nextGroupStart
		prevGroupEnd = groupStart
	}
	return dummy.Next    
}

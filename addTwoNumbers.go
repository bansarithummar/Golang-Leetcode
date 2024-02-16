2. Add Two Numbers

package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Initialize dummy head node for the result list
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	// Iterate through both linked lists until both are nil and carry is zero
	for l1 != nil || l2 != nil || carry > 0 {
		// Calculate the sum of the current digits and carry
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Update carry and create a new node with the sum digit
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	// Return the next node after the dummy head, which contains the result
	return dummyHead.Next
}

// Helper function to create a linked list from an array of integers
func createList(nums []int) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummyHead.Next
}

// Helper function to print the linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}




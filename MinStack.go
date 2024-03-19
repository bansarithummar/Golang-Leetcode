155. Min Stack

package main

import "fmt"

type MinStack struct {
    stack    []int // Stack to store all elements
    minStack []int // Stack to store the minimum elements
}

func Constructor() MinStack {
    return MinStack{
        stack:    []int{},
        minStack: []int{},
    }
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)
    if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, val)
    }
}

func (this *MinStack) Pop() {
    if len(this.stack) > 0 {
        topElement := this.stack[len(this.stack)-1]
        this.stack = this.stack[:len(this.stack)-1]
        if topElement == this.minStack[len(this.minStack)-1] {
            this.minStack = this.minStack[:len(this.minStack)-1]
        }
    }
}

func (this *MinStack) Top() int {
    if len(this.stack) > 0 {
        return this.stack[len(this.stack)-1]
    }
    return -1 // Or handle error appropriately
}

func (this *MinStack) GetMin() int {
    if len(this.minStack) > 0 {
        return this.minStack[len(this.minStack)-1]
    }
    return -1 // Or handle error appropriately
}

// Ensure to integrate the testing part in your own main function, or run this snippet in isolation.



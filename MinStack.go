155. Min Stack

type MinStack struct {
    stack    []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{
        stack:    make([]int, 0),
        minStack: make([]int, 0),
    }
}

func (s *MinStack) Push(val int) {
    s.stack = append(s.stack, val)
    minVal := val
    if len(s.minStack) > 0 && s.minStack[len(s.minStack)-1] < val {
        minVal = s.minStack[len(s.minStack)-1]
    }
    s.minStack = append(s.minStack, minVal)
}

func (s *MinStack) Pop() {
    if len(s.stack) > 0 {
        s.stack = s.stack[:len(s.stack)-1]
        s.minStack = s.minStack[:len(s.minStack)-1]
    }
}

func (s *MinStack) Top() int {
    if len(s.stack) > 0 {
        return s.stack[len(s.stack)-1]
    }
    return -1
}

func (s *MinStack) GetMin() int {
    if len(s.minStack) > 0 {
        return s.minStack[len(s.minStack)-1]
    }
    return math.MaxInt64
}


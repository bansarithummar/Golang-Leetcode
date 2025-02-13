func minOperations(nums []int, k int) int {
    h := &hp{} 
   heap.Init(h)
   
   // Add all nums to heap
   for _, num := range nums {
       heap.Push(h, num)
   }
   
   operations := 0
   
   // Process until min element >= k
   for h.Len() >= 2 && h.IntSlice[0] < k {
       // Get two smallest elements
       x := heap.Pop(h).(int)
       y := heap.Pop(h).(int)
       
       // Calculate new element and add back
       newVal := min(x,y)*2 + max(x,y)
       heap.Push(h, newVal)
       operations++
   }
   
   return operations
}

// Heap helper struct for built-in heap interface
type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
   n := len(h.IntSlice)
   v := h.IntSlice[n-1]
   h.IntSlice = h.IntSlice[:n-1]
   return v
}

func min(a, b int) int {
   if a < b {
       return a 
   }
   return b
}

func max(a, b int) int {
   if a > b {
       return a
   }
   return b   
}

func punishmentNumber(n int) int {
    result := 0
   
   for i := 1; i <= n; i++ {
       square := i * i
       // If valid partition exists, add to result
       if canPartition(square, i) {
           result += square
       }
   }
   return result
}

// Helper to check if num can be partitioned into parts summing to target
func canPartition(num int, target int) bool {
   // Convert number to string for easier partitioning
   numStr := strconv.Itoa(num)
   
   // Try all possible partitions using backtracking
   var backtrack func(idx, sum int) bool
   backtrack = func(idx, sum int) bool {
       // Base case: reached end of string
       if idx == len(numStr) {
           return sum == target
       }
       
       // Try partitions starting at current index
       for i := idx; i < len(numStr); i++ {
           // Get current partition value
           val, _ := strconv.Atoi(numStr[idx:i+1])
           if sum + val <= target && backtrack(i+1, sum + val) {
               return true
           }
       }
       return false
   }  
   return backtrack(0, 0)
}

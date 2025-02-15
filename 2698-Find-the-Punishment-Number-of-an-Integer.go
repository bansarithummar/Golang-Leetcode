func punishmentNumber(n int) int {
    result := 0
   
   for i := 1; i <= n; i++ {
       square := i * i
       if canPartition(square, i) {
           result += square
       }
   }
   return result
}

func canPartition(num int, target int) bool {
   numStr := strconv.Itoa(num)
   var backtrack func(idx, sum int) bool
   backtrack = func(idx, sum int) bool {
       if idx == len(numStr) {
           return sum == target
       }

       for i := idx; i < len(numStr); i++ {
           val, _ := strconv.Atoi(numStr[idx:i+1])
           if sum + val <= target && backtrack(i+1, sum + val) {
               return true
           }
       }
       return false
   }  
   return backtrack(0, 0)
}

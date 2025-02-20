func findDifferentBinaryString(nums []string) string {
    n := len(nums)
   result := make([]byte, n)
  
   for i := 0; i < n; i++ {
       if nums[i][i] == '0' {
           result[i] = '1'
       } else {
           result[i] = '0'
       }
   }
   
   return string(result) 
}

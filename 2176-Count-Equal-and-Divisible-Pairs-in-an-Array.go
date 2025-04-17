
func countPairs(nums []int, k int) int {
    count := 0
    n := len(nums)
    
    // Check all possible pairs (i,j) where i < j
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            // Check if nums[i] == nums[j] and (i*j) is divisible by k
            if nums[i] == nums[j] && (i*j)%k == 0 {
                count++
            }
        }
    }
    
    return count
    
}

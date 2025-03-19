func minOperations(nums []int) int {
    n := len(nums)
    count := 0
    temp := make([]int, n)
    copy(temp, nums)
    
    for i := 0; i < n-2; i++ {
        if temp[i] == 0 {
            temp[i] ^= 1
            temp[i+1] ^= 1
            temp[i+2] ^= 1
            count++
        }
    }
    
    for i := 0; i < n; i++ {
        if temp[i] == 0 {
            return -1
        }
    }  
    return count  
}

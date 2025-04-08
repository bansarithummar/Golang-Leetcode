func minimumOperations(nums []int) int {
    n := len(nums)
    
    for ops := 0; ops <= (n+2)/3; ops++ {
        remaining := nums
        if 3*ops < n {
            remaining = nums[3*ops:]
        } else {
            remaining = []int{}
        }
        
        if hasDistinctElements(remaining) {
            return ops
        }
    }  
    return (n+2)/3 
}

func hasDistinctElements(nums []int) bool {
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] {
            return false
        }
        seen[num] = true
    }
    return true
}

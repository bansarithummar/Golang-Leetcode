func jump(nums []int) int {
    jumps := 0
    current_end := 0
    furthest := 0
    for i := 0; i < len(nums) - 1; i++ {
        if i + nums[i] > furthest {
            furthest = i + nums[i]
        }
        if i == current_end {
            jumps++
            current_end = furthest
            if current_end >= len(nums) - 1 {
                break
            }
        }
    }
    return jumps    
}

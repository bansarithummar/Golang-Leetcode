func canJump(nums []int) bool {
    furthest := 0
    for i := 0; i < len(nums); i++ {
        if i > furthest {
            return false
        }
        furthest = max(furthest, i + nums[i])
        if furthest >= len(nums) - 1 {
            return true
        }
    }
    return false
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b   
}

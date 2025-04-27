func countSubarrays(nums []int) int {
    count := 0
    for i := 0; i+2 < len(nums); i++ {
        mid := nums[i+1]
        if mid%2 == 0 && nums[i]+nums[i+2] == mid/2 {
            count++
        }
    }
    return count
}

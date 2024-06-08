func findDuplicate(nums []int) int {
    tortoise := nums[0]
    hare := nums[0]
    for {
        tortoise = nums[tortoise]
        hare = nums[nums[hare]]
        if tortoise == hare {
            break
        }
    }
    ptr1 := nums[0]
    ptr2 := tortoise
    for ptr1 != ptr2 {
        ptr1 = nums[ptr1]
        ptr2 = nums[ptr2]
    }
    
    return ptr1    
}

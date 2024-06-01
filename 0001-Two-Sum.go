func twoSum(nums []int, target int) []int {
    prevMap := make(map[int]int)

    for i, num := range nums {
        diff := target - num
        if j, found := prevMap[diff]; found {
            return []int{j, i}
        }
        prevMap[num] = i
    }
    return nil    
}

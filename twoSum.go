1. Two Sum

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)

    // Iterate over the array
    for i, num := range nums {
        complement := target - num
        if j, ok := numMap[complement]; ok {
            // If the complement is found in the map, return the indices
            return []int{j, i}
        }
        // Store the index of the current number in the map
        numMap[num] = i
    }
    // In case no solution is found, though the problem guarantees one.
    return nil    
}

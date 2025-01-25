func lexicographicallySmallestArray(nums []int, limit int) []int {
    n := len(nums)
    indices := make([]int, n)
    for i := range indices {
        indices[i] = i
    }
    
    sort.Slice(indices, func(i, j int) bool {
        return nums[indices[i]] < nums[indices[j]]
    })
    
    groups := [][]int{}
    currGroup := []int{indices[0]}
    
    for i := 1; i < n; i++ {
        // Check if current element can be added to the group
        if nums[indices[i]] - nums[indices[i-1]] <= limit {
            currGroup = append(currGroup, indices[i])
        } else {
            // Start a new group
            groups = append(groups, currGroup)
            currGroup = []int{indices[i]}
        }
    }
    groups = append(groups, currGroup)
    
    // Sort indices within each group
    result := make([]int, n)
    for _, group := range groups {
        // Sort original values within the group
        sort.Slice(group, func(i, j int) bool {
            return group[i] < group[j]
        })
        
        // Sort group values
        sortedGroupValues := make([]int, len(group))
        for i, idx := range group {
            sortedGroupValues[i] = nums[idx]
        }
        sort.Ints(sortedGroupValues)
        
        // Assign sorted values back to original positions
        for i, idx := range group {
            result[idx] = sortedGroupValues[i]
        }
    }
    
    return result
    
}

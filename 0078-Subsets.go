func subsets(nums []int) [][]int {
    var result [][]int
    var current []int
    backtrack(nums, 0, current, &result)
    return result
}

func backtrack(nums []int, start int, current []int, result *[][]int) {
    subset := make([]int, len(current))
    copy(subset, current)
    *result = append(*result, subset)
    for i := start; i < len(nums); i++ {
        current = append(current, nums[i])
        backtrack(nums, i+1, current, result)
        current = current[:len(current)-1] 
    }    
}

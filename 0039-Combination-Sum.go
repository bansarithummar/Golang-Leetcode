func combinationSum(candidates []int, target int) [][]int {
    var result [][]int
    var current []int
    backtrack(candidates, target, 0, current, &result)
    return result
}


func backtrack(candidates []int, target int, start int, current []int, result *[][]int) {
    if target == 0 {
        combination := make([]int, len(current))
        copy(combination, current)
        *result = append(*result, combination)
        return
    }
    
    for i := start; i < len(candidates); i++ {
        if candidates[i] <= target {
            current = append(current, candidates[i])
            backtrack(candidates, target - candidates[i], i, current, result)
            current = current[:len(current)-1]
        }
    }
}

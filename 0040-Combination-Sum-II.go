func combinationSum2(candidates []int, target int) [][]int {
    var result [][]int
    var current []int
    sort.Ints(candidates)
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
        if i > start && candidates[i] == candidates[i-1] {
            continue
        }
        if candidates[i] > target {
            break
        }
        current = append(current, candidates[i])
        backtrack(candidates, target-candidates[i], i+1, current, result)
        current = current[:len(current)-1]
    }    
}

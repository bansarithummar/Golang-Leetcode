2149. Rearrange Array Elements by Sign

func rearrangeArray(nums []int) []int {
    pi := 0
    ni := 1
    res := make([]int, len(nums))

    for _, n := range nums {
        if n > 0 {
            res[pi] = n
            pi += 2
        } else {
            res[ni] = n
            ni += 2
        }
    }

    return res
}

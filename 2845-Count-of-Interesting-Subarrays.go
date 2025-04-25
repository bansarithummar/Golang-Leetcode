func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
    count := int64(0)
    prefixModCount := map[int]int{}
    prefixModCount[0] = 1 

    curr := 0
    for _, num := range nums {
        if num%modulo == k {
            curr++
        }
        currMod := curr % modulo
        want := (currMod - k + modulo) % modulo
        count += int64(prefixModCount[want])
        prefixModCount[currMod]++
    }

    return count
    
}

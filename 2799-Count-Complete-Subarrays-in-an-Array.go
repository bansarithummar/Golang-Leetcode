func countCompleteSubarrays(nums []int) int {
    distinctSet := make(map[int]bool)
    for _, num := range nums {
        distinctSet[num] = true
    }
    totalDistinct := len(distinctSet)

    count := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        freq := make(map[int]int)
        distinct := 0
        for j := i; j < n; j++ {
            if freq[nums[j]] == 0 {
                distinct++
            }
            freq[nums[j]]++
            if distinct == totalDistinct {
                count++
            }
        }
    }

    return count
}

func topKFrequent(nums []int, k int) []int {
    frequencyMap := make(map[int]int)
    for _, num := range nums {
        frequencyMap[num]++
    }
    buckets := make([][]int, len(nums)+1)
    for num, freq := range frequencyMap {
        buckets[freq] = append(buckets[freq], num)
    }
    result := make([]int, 0, k)
    for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
        if len(buckets[i]) > 0 {
            result = append(result, buckets[i]...)
        }
    }
    return result[:k]
}

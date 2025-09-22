func maxFrequencyElements(nums []int) int {
    freq := make(map[int]int)
    maxFreq := 0
    for _, v := range nums {
        freq[v]++
        if freq[v] > maxFreq {
            maxFreq = freq[v]
        }
    }
    res := 0
    for _, f := range freq {
        if f == maxFreq {
            res += f
        }
    }
    return res
}

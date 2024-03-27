1481. Least Number of Unique Integers after K Removals

func findLeastNumOfUniqueInts(arr []int, k int) int {
    freqMap := make(map[int]int)
    for _, num := range arr {
        freqMap[num]++
    }
    sortedFreq := make([][2]int, 0, len(freqMap))
    for num, freq := range freqMap {
        sortedFreq = append(sortedFreq, [2]int{num, freq})
    }
    sort.Slice(sortedFreq, func(i, j int) bool {
        return sortedFreq[i][1] < sortedFreq[j][1]
    })
    for _, pair := range sortedFreq {
        if k >= pair[1] {
            k -= pair[1]
            delete(freqMap, pair[0])
        } else {
            break
        }
    }
    return len(freqMap)
}

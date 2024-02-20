1481. Least Number of Unique Integers after K Removals

import (
    "sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
    freqMap := make(map[int]int)
    for _, num := range arr {
        freqMap[num]++
    }

    // Convert the frequency map to a slice of pairs (num, freq)
    sortedFreq := make([][2]int, 0, len(freqMap))
    for num, freq := range freqMap {
        sortedFreq = append(sortedFreq, [2]int{num, freq})
    }

    // Sort the slice based on frequency
    sort.Slice(sortedFreq, func(i, j int) bool {
        return sortedFreq[i][1] < sortedFreq[j][1]
    })

    // Remove elements based on k until k becomes 0 or there are no more elements to remove
    for _, pair := range sortedFreq {
        if k >= pair[1] {
            k -= pair[1]
            delete(freqMap, pair[0])
        } else {
            break
        }
    }

    // The remaining elements in freqMap represent the least number of unique integers
    return len(freqMap)
}

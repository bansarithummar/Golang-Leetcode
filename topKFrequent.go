347. Top K Frequent Elements

package main

import "sort"

func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    for _, n := range nums {
        count[n]++
    }
    
    freq := make([][]int, len(nums)+1)
    for n, c := range count {
        freq[c] = append(freq[c], n)
    }

    var res []int
    for i := len(freq) - 1; i > 0; i-- {
        for _, n := range freq[i] {
            res = append(res, n)
            if len(res) == k {
                return res
            }
        }
    }
    return res
}

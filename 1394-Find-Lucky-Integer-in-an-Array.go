func findLucky(arr []int) int {
    freqMap := make(map[int]int)
    
    for _, num := range arr {
        freqMap[num]++
    }

    res := -1
    for num, freq := range freqMap {
        if num == freq && num > res {
            res = num
        }
    }

    return res   
}

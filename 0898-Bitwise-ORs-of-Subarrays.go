func subarrayBitwiseORs(arr []int) int {
    prevSet := map[int]bool{}
    resSet := map[int]bool{}

    for _, num := range arr {
        currSet := map[int]bool{}
        currSet[num] = true
        for val := range prevSet {
            currSet[val|num] = true
        }
        for val := range currSet {
            resSet[val] = true
        }
        prevSet = currSet
    }

    return len(resSet)
    
}

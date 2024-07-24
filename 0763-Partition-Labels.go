func partitionLabels(s string) []int {
    lastOccurrence := make(map[rune]int)
    for i, char := range s {
        lastOccurrence[char] = i
    }    
    var partitions []int
    start, end := 0, 0
    for i, char := range s {
        if lastOccurrence[char] > end {
            end = lastOccurrence[char]
        }
        if i == end {
            partitions = append(partitions, end - start + 1)
            start = i + 1
        }
    }   
    return partitions   
}

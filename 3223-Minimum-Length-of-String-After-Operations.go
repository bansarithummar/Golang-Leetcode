func minimumLength(s string) int {
    freqMap := make(map[rune]int)

    for _, c := range s {
        freqMap[c]++
    }

    deletions := 0

    for _, count := range freqMap {
        if count > 2 {
            if count%2 == 0 {
                deletions += (count - 2)
            } else {
                deletions += (count - 1)
            }
        }
    }

    return len(s) - deletions  
}

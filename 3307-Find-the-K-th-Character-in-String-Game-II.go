func kthCharacter(k int64, operations []int) byte {
    var shifts int64
    for k != 1 {
        i := 63 - bits.LeadingZeros64(uint64(k))
        half := int64(1) << i
        if half == k {
            i--
            half = int64(1) << i
        }
        k -= half
        if operations[i] == 1 {
            shifts++
        }
    }
    return byte(shifts%26 + 97)  
}

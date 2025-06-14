func minMaxDifference(num int) int {
    s := strconv.Itoa(num)

    maxRunes := []rune(s)
    var r rune
    for _, ch := range maxRunes {
        if ch != '9' {
            r = ch
            break
        }
    }
    if r != 0 {
        for i, ch := range maxRunes {
            if ch == r {
                maxRunes[i] = '9'
            }
        }
    }
    maxVal, _ := strconv.Atoi(string(maxRunes))

    minRunes := []rune(s)
    r = 0
    for _, ch := range minRunes {
        if ch != '0' {
            r = ch
            break
        }
    }
    if r != 0 {
        for i, ch := range minRunes {
            if ch == r {
                minRunes[i] = '0'
            }
        }
    }
    minVal, _ := strconv.Atoi(string(minRunes))

    return maxVal - minVal   
}

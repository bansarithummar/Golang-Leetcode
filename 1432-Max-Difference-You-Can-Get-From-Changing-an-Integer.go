func maxDiff(num int) int {
    s := strconv.Itoa(num)

    maxBytes := []byte(s)
    var tgt byte
    for _, b := range maxBytes {
        if b != '9' {
            tgt = b
            break
        }
    }
    if tgt != 0 {
        for i, b := range maxBytes {
            if b == tgt {
                maxBytes[i] = '9'
            }
        }
    }
    maxVal, _ := strconv.Atoi(string(maxBytes))

    minBytes := []byte(s)
    if minBytes[0] != '1' {
        tgt = minBytes[0]
        for i, b := range minBytes {
            if b == tgt {
                minBytes[i] = '1'
            }
        }
    } else {
        tgt = 0
        for _, b := range minBytes[1:] {
            if b != '0' && b != '1' {
                tgt = b
                break
            }
        }
        if tgt != 0 {
            for i, b := range minBytes {
                if b == tgt {
                    minBytes[i] = '0'
                }
            }
        }
    }
    minVal, _ := strconv.Atoi(string(minBytes))

    return maxVal - minVal
}

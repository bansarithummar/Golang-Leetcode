func makeFancyString(s string) string {
    var res []byte
    count := 1

    for i := 0; i < len(s); i++ {
        if i > 0 && s[i] == s[i-1] {
            count++
        } else {
            count = 1
        }

        if count < 3 {
            res = append(res, s[i])
        }
    }

    return string(res)
    
}

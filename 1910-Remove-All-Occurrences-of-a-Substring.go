func removeOccurrences(s string, part string) string {
    if len(s) == 0 || len(part) == 0 {
        return s
    }
    var builder strings.Builder

    for i := 0; i < len(s); i++ {
        builder.WriteByte(s[i])

        currLen := builder.Len()
        if currLen >= len(part) {

            curr := builder.String()
            if curr[currLen-len(part):] == part {
                builder.Reset()
                builder.WriteString(curr[:currLen-len(part)])
            }
        }
    }    
    return builder.String()
}

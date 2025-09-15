func canBeTypedWords(text string, brokenLetters string) int {
    broken := make(map[rune]bool)
    for _, ch := range brokenLetters {
        broken[ch] = true
    }
    count := 0
    words := strings.Split(text, " ")
    for _, w := range words {
        ok := true
        for _, ch := range w {
            if broken[ch] {
                ok = false
                break
            }
        }
        if ok {
            count++
        }
    }
    return count   
}

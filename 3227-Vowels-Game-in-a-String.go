func doesAliceWin(s string) bool {
    vowels := map[byte]bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    }
    count := 0
    for i := 0; i < len(s); i++ {
        if vowels[s[i]] {
            count++
        }
    }
    // Alice wins if there is at least one vowel (odd substring exists)
    return count > 0
    
}

func reorderedPowerOf2(n int) bool {
    target := countDigits(n)
    for i := 0; i < 31; i++ {
        if equalCounts(target, countDigits(1<<i)) {
            return true
        }
    }
    return false
}

func countDigits(num int) [10]int {
    var cnt [10]int
    for _, ch := range strconv.Itoa(num) {
        cnt[ch-'0']++
    }
    return cnt
}

func equalCounts(a, b [10]int) bool {
    for i := 0; i < 10; i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true   
}

func partition(s string) [][]string {
    var result [][]string
    var current []string
    backtrack(s, 0, &current, &result)
    return result
}


func backtrack(s string, start int, current *[]string, result *[][]string) {
    if start == len(s) {
        partition := make([]string, len(*current))
        copy(partition, *current)
        *result = append(*result, partition)
        return
    }  
    for end := start; end < len(s); end++ {
        if isPalindrome(s, start, end) {
            *current = append(*current, s[start:end+1])
            backtrack(s, end+1, current, result)
            *current = (*current)[:len(*current)-1]
        }
    }   
}


func isPalindrome(s string, left, right int) bool {
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}


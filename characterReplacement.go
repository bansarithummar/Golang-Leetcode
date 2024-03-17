424. Longest Repeating Character Replacement

func characterReplacement(s string, k int) int {
    maxLen := 0
    maxCount := 0
    count := make(map[byte]int)
    left := 0

    for right := 0; right < len(s); right++ {
        count[s[right]]++
        maxCount = max(maxCount, count[s[right]])

        // If the current window size minus the frequency of the most frequent character
        // is greater than k, shrink the window from the left.
        if right-left+1-maxCount > k {
            count[s[left]]--
            left++
        }

        // Update maxLen if the current window size is larger.
        maxLen = max(maxLen, right-left+1)
    }

    return maxLen
}

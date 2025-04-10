func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
    for _, ch := range s {
        if int(ch-'0') > limit {
            return 0
        }
    }
    
    // Function to count powerful integers up to n
    countUpTo := func(n int64) int64 {
        nStr := strconv.FormatInt(n, 10)
        suffixLen := len(s)
        
        // If n's length is less than suffix length, no powerful integers
        if len(nStr) < suffixLen {
            return 0
        }
        
        // Memoization map for dynamic programming
        memo := make(map[string]int64)
        
        // Recursive DP function
        var dp func(int, bool) int64
        dp = func(pos int, tight bool) int64 {
            // Base case: finished processing all positions
            if pos == len(nStr) {
                return 1
            }
            
            // Memoization key
            key := fmt.Sprintf("%d_%t", pos, tight)
            if val, ok := memo[key]; ok {
                return val
            }
            
            // Check if we're in the suffix part
            inSuffix := pos >= len(nStr) - suffixLen
            var result int64 = 0
            
            // Determine upper bound for current position
            upperBound := limit
            if tight {
                upperBound = min(upperBound, int(nStr[pos]-'0'))
            }
            
            // If in suffix part, we must match the corresponding suffix digit
            if inSuffix {
                suffixPos := pos - (len(nStr) - suffixLen)
                targetDigit := int(s[suffixPos] - '0')
                
                // If the required suffix digit is valid, continue
                if targetDigit <= upperBound {
                    newTight := tight && (targetDigit == int(nStr[pos]-'0'))
                    result += dp(pos+1, newTight)
                }
            } else {
                // If not in suffix, try all valid digits
                for digit := 0; digit <= upperBound; digit++ {
                    newTight := tight && (digit == int(nStr[pos]-'0'))
                    result += dp(pos+1, newTight)
                }
            }
            
            memo[key] = result
            return result
        }
        
        return dp(0, true)
    }
    
    // Result: count up to finish minus count up to (start-1)
    return countUpTo(finish) - countUpTo(start-1)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b  
}

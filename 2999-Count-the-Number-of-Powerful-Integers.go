func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
    for _, ch := range s {
        if int(ch-'0') > limit {
            return 0
        }
    }
    
    countUpTo := func(n int64) int64 {
        nStr := strconv.FormatInt(n, 10)
        suffixLen := len(s)
        
        if len(nStr) < suffixLen {
            return 0
        }
        
        memo := make(map[string]int64)
        
        var dp func(int, bool) int64
        dp = func(pos int, tight bool) int64 {
            
            if pos == len(nStr) {
                return 1
            }
            
            key := fmt.Sprintf("%d_%t", pos, tight)
            if val, ok := memo[key]; ok {
                return val
            }
            
            inSuffix := pos >= len(nStr) - suffixLen
            var result int64 = 0

            upperBound := limit
            if tight {
                upperBound = min(upperBound, int(nStr[pos]-'0'))
            }
            
            if inSuffix {
                suffixPos := pos - (len(nStr) - suffixLen)
                targetDigit := int(s[suffixPos] - '0')

                if targetDigit <= upperBound {
                    newTight := tight && (targetDigit == int(nStr[pos]-'0'))
                    result += dp(pos+1, newTight)
                }
            } else {
                
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
    return countUpTo(finish) - countUpTo(start-1)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b  
}

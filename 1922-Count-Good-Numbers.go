func countGoodNumbers(n int64) int {
    mod := int64(1e9 + 7)
    
    // For even positions: 0, 2, 4, 6, 8 (5 options)
    // For odd positions: 2, 3, 5, 7 (4 options)
  
    evenPositions := (n + 1) / 2 // Ceiling division for odd n
    oddPositions := n / 2        // Floor division
    
    return int(modPow(5, evenPositions, mod) * modPow(4, oddPositions, mod) % mod)
}

func modPow(base, exp, mod int64) int64 {
    if exp == 0 {
        return 1
    }
    
    result := int64(1)
    base %= mod
    
    for exp > 0 {
        if exp % 2 == 1 {
            result = (result * base) % mod
        }
        
        base = (base * base) % mod
        
        exp /= 2
    }
    
    return result 
}

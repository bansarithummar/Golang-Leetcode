func countGoodNumbers(n int64) int {
    mod := int64(1e9 + 7)

    evenPositions := (n + 1) / 2 
    oddPositions := n / 2     
    
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

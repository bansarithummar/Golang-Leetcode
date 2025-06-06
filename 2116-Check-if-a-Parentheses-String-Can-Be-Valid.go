func canBeValid(s string, locked string) bool {
    n := len(s)
    if n%2 != 0 {
        return false
    }

    openCount, unlockCount := 0, 0
    for i := 0; i < n; i++ {
        if locked[i] == '0' {
            unlockCount++
        } else if s[i] == '(' {
            openCount++
        } else { 
            if openCount > 0 {
                openCount--
            } else if unlockCount > 0 {
                unlockCount--
            } else {
                return false
            }
        }
    }

    closeCount, unlockCount := 0, 0
    for i := n - 1; i >= 0; i-- {
        if locked[i] == '0' {
            unlockCount++
        } else if s[i] == ')' {
            closeCount++
        } else { 
            if closeCount > 0 {
                closeCount--
            } else if unlockCount > 0 {
                unlockCount--
            } else {
                return false
            }
        }
    }

    return true
}

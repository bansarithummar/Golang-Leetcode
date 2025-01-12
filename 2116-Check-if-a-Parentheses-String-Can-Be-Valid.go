func canBeValid(s string, locked string) bool {
    n := len(s)
    // If the length of the string is odd, it cannot be valid
    if n%2 != 0 {
        return false
    }

    // Check from left to right
    openCount, unlockCount := 0, 0
    for i := 0; i < n; i++ {
        if locked[i] == '0' {
            unlockCount++
        } else if s[i] == '(' {
            openCount++
        } else { // s[i] == ')' and locked[i] == '1'
            if openCount > 0 {
                openCount--
            } else if unlockCount > 0 {
                unlockCount--
            } else {
                return false
            }
        }
    }

    // Check from right to left
    closeCount, unlockCount := 0, 0
    for i := n - 1; i >= 0; i-- {
        if locked[i] == '0' {
            unlockCount++
        } else if s[i] == ')' {
            closeCount++
        } else { // s[i] == '(' and locked[i] == '1'
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

func pushDominoes(dominoes string) string {
    n := len(dominoes)
    forces := make([]int, n)
    
    // Apply force from 'R'
    force := 0
    for i := 0; i < n; i++ {
        if dominoes[i] == 'R' {
            force = n
        } else if dominoes[i] == 'L' {
            force = 0
        } else if force > 0 {
            force--
        }
        forces[i] += force
    }

    // Apply force from 'L'
    force = 0
    for i := n - 1; i >= 0; i-- {
        if dominoes[i] == 'L' {
            force = n
        } else if dominoes[i] == 'R' {
            force = 0
        } else if force > 0 {
            force--
        }
        forces[i] -= force
    }

    // Build the result
    result := make([]byte, n)
    for i := 0; i < n; i++ {
        if forces[i] > 0 {
            result[i] = 'R'
        } else if forces[i] < 0 {
            result[i] = 'L'
        } else {
            result[i] = '.'
        }
    }

    return string(result)
    
}

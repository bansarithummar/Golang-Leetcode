func maxCollectedFruits(fruits [][]int) int {
    n := len(fruits)
    total := 0
    for i := 0; i < n; i++ {
        total += fruits[i][i]
    }

    rightPath := make([]int, 3)
    rightPath[0] = fruits[0][n-1]

    bottomPath := make([]int, 3)
    bottomPath[0] = fruits[n-1][0]

    window := 2

    for step := 1; step < n-1; step++ {
        newRight := make([]int, window+2)
        newBottom := make([]int, window+2)

        for dist := 0; dist < window; dist++ {
            left, mid, right := 0, rightPath[dist], 0
            if dist-1 >= 0 {
                left = rightPath[dist-1]
            }
            if dist+1 < len(rightPath) {
                right = rightPath[dist+1]
            }
            newRight[dist] = max(left, max(mid, right)) + fruits[step][n-1-dist]

            left, mid, right = 0, bottomPath[dist], 0
            if dist-1 >= 0 {
                left = bottomPath[dist-1]
            }
            if dist+1 < len(bottomPath) {
                right = bottomPath[dist+1]
            }
            newBottom[dist] = max(left, max(mid, right)) + fruits[n-1-dist][step]
        }

        rightPath = newRight
        bottomPath = newBottom

        if window-n+4+step <= 1 {
            window++
        } else if window-n+3+step > 1 {
            window--
        }
    }

    return total + rightPath[0] + bottomPath[0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b   
}

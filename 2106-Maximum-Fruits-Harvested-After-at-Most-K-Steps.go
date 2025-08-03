func maxTotalFruits(fruits [][]int, startPos int, k int) int {
    n := len(fruits)
    prefixSum := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefixSum[i+1] = prefixSum[i] + fruits[i][1]
    }

    maxFruits := 0

    left := 0
    for right := 0; right < n; right++ {
        for left <= right {
            l := fruits[left][0]
            r := fruits[right][0]
            dist := min(abs(startPos-l)+r-l, abs(startPos-r)+r-l)
            if dist <= k {
                break
            }
            left++
        }
        if left <= right {
            sum := prefixSum[right+1] - prefixSum[left]
            if sum > maxFruits {
                maxFruits = sum
            }
        }
    }

    return maxFruits
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b    
}

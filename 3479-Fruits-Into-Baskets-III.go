func numOfUnplacedFruits(fruits []int, baskets []int) int {
    n := len(baskets)
    N := 1
    for N <= n {
        N <<= 1
    }

    segTree := make([]int, N<<1)
    for i := 0; i < n; i++ {
        segTree[N+i] = baskets[i]
    }

    for i := N - 1; i > 0; i-- {
        segTree[i] = max(segTree[2*i], segTree[2*i+1])
    }

    count := 0
    for _, fruit := range fruits {
        index := 1
        if segTree[index] < fruit {
            count++
            continue
        }

        for index < N {
            if segTree[2*index] >= fruit {
                index = 2 * index
            } else {
                index = 2*index + 1
            }
        }

        segTree[index] = -1
        for index > 1 {
            index >>= 1
            segTree[index] = max(segTree[2*index], segTree[2*index+1])
        }
    }

    return count
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b   
}

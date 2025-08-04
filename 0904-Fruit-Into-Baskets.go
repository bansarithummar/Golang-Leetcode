func totalFruit(fruits []int) int {
    basket := make(map[int]int)
    left := 0
    maxFruits := 0

    for right, fruit := range fruits {
        basket[fruit]++

        for len(basket) > 2 {
            basket[fruits[left]]--
            if basket[fruits[left]] == 0 {
                delete(basket, fruits[left])
            }
            left++
        }

        if right-left+1 > maxFruits {
            maxFruits = right - left + 1
        }
    }
    return maxFruits  
}

func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand) % groupSize != 0 {
        return false
    }
    cardCount := make(map[int]int)
    for _, card := range hand {
        cardCount[card]++
    }
    sort.Ints(hand)
    for _, card := range hand {
        if cardCount[card] > 0 {
            for i := 0; i < groupSize; i++ {
                if cardCount[card + i] == 0 {
                    return false
                }
                cardCount[card + i]--
            }
        }
    }
    return true
}

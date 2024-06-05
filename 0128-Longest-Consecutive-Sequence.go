func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }
    longestStreak := 0
    for num := range numSet {
        if !numSet[num-1] {
            currentNum := num
            currentStreak := 1
            for numSet[currentNum+1] {
                currentNum += 1
                currentStreak += 1
            }
            if currentStreak > longestStreak {
                longestStreak = currentStreak
            }
        }
    }
    return longestStreak
}

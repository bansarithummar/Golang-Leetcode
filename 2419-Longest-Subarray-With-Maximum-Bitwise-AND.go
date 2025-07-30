func longestSubarray(nums []int) int {
    maxVal := 0
    for _, num := range nums {
        if num > maxVal {
            maxVal = num
        }
    }

    longest := 0
    count := 0

    for _, num := range nums {
        if num == maxVal {
            count++
            if count > longest {
                longest = count
            }
        } else {
            count = 0
        }
    }

    return longest   
}

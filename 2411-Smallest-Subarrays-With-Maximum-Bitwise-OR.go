func smallestSubarrays(nums []int) []int {
    n := len(nums)
    answer := make([]int, n)
    lastSeen := make([]int, 32) // last index for each bit

    for i := 0; i < 32; i++ {
        lastSeen[i] = -1
    }

    for i := n - 1; i >= 0; i-- {
        for bit := 0; bit < 32; bit++ {
            if (nums[i]>>bit)&1 == 1 {
                lastSeen[bit] = i
            }
        }

        maxIdx := i
        for bit := 0; bit < 32; bit++ {
            if lastSeen[bit] > maxIdx {
                maxIdx = lastSeen[bit]
            }
        }

        answer[i] = maxIdx - i + 1
    }

    return answer   
}

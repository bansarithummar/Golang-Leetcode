func minimizeMax(nums []int, p int) int {
    if p == 0 {
        return 0
    }

    sort.Ints(nums)
    low, high := 0, nums[len(nums)-1]-nums[0]

    canMake := func(maxDiff int) bool {
        pairs := 0
        for i := 1; i < len(nums) && pairs < p; {
            if nums[i]-nums[i-1] <= maxDiff {
                pairs++
                i += 2
            } else {
                i++
            }
        }
        return pairs >= p
    }

    for low < high {
        mid := (low + high) / 2
        if canMake(mid) {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return low
}

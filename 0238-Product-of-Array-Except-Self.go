func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    prefix := 1
    for i := 0; i < n; i++ {
        res[i] = prefix
        prefix *= nums[i]
    }
    postfix := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }

    return res
}

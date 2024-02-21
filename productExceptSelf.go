238. Product of Array Except Self

package main

import "fmt"

func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    // Calculate prefix products
    prefix := 1
    for i := 0; i < n; i++ {
        res[i] = prefix
        prefix *= nums[i]
    }

    // Calculate postfix products while updating the result
    postfix := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }

    return res
}

func main() {
    nums := []int{1, 2, 3, 4}
    result := productExceptSelf(nums)
    fmt.Println(result) // Output: [24 12 8 6]
}

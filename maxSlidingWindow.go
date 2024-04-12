239. Sliding Window Maximum

func maxSlidingWindow(nums []int, k int) []int {
    result := make([]int, 0)
    deque := make([]int, 0)
    add := func(i int) {
        for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
    }
    for i := 0; i < k; i++ {
        add(i)
    }
    result = append(result, nums[deque[0]])
    for i := k; i < len(nums); i++ {
        if deque[0] == i-k {
            deque = deque[1:]
        }
        add(i)
        result = append(result, nums[deque[0]])
    }
    return result    
}

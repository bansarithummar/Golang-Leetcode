739. Daily Temperatures


func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    ans := make([]int, n)

    var stack []int
    for i, temp := range temperatures {
        for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
            index := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans[index] = i - index
        }
        stack = append(stack, i)
    }
    
    return ans 
}


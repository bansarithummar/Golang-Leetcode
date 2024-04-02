1. Two Sum

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if j, ok := numMap[complement]; o                                         
            return []int{j, i}
        }
        numMap[num] = i
    }
    return nil    
}

func minimumIndex(nums []int) int {
    dominant := findDominantElement(nums)
    
    totalCount := 0
    for _, num := range nums {
        if num == dominant {
            totalCount++
        }
    }

    leftCount, rightCount := 0, totalCount
    
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == dominant {
            leftCount++
            rightCount--
        }
        
        if leftCount * 2 > i+1 && rightCount * 2 > len(nums)-i-1 {
            return i
        }
    }
    
    return -1
}

func findDominantElement(nums []int) int {
    candidate, count := 0, 0
    
    for _, num := range nums {
        if count == 0 {
            candidate = num
        }
        
        if num == candidate {
            count++
        } else {
            count--
        }
    }
    
    return candidate
}

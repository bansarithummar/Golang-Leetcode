func maximumTripletValue(nums []int) int64 {
    maxNum := 0
    maxDiff := 0
    result := int64(0)
    
    for _, num := range nums {
        curValue := int64(maxDiff * num)
        
        if curValue > result {
            result = curValue
        }
        
        if maxNum - num > maxDiff {
            maxDiff = maxNum - num
        }

        if num > maxNum {
            maxNum = num
        }
    }
    
    return result   
}

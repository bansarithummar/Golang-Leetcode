func largestGoodInteger(num string) string {
    res := ""
    for i := 0; i <= len(num)-3; i++ {
        if num[i] == num[i+1] && num[i+1] == num[i+2] {
            cur := num[i : i+3]
            if cur > res {
                res = cur
            }
        }
    }
    return res
    
}

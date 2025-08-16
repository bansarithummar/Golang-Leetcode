func maximum69Number (num int) int {
    digits := []byte(strconv.Itoa(num))
    for i := 0; i < len(digits); i++ {
        if digits[i] == '6' {
            digits[i] = '9'
            break
        }
    }
    ans, _ := strconv.Atoi(string(digits))
    return ans   
}

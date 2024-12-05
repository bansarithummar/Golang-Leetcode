func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    m, n := len(num1), len(num2)
    result := make([]int, m+n)

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            mul := int(num1[i]-'0') * int(num2[j]-'0')
            sum := mul + result[i+j+1]

            result[i+j+1] = sum % 10
            result[i+j] += sum / 10
        }
    }

    resStr := ""
    for _, r := range result {
        if !(len(resStr) == 0 && r == 0) { 
            resStr += string(r + '0')
        }
    }
    return resStr  
}

func replaceNonCoprimes(nums []int) []int {
    gcd := func(a, b int) int {
        for b != 0 {
            a, b = b, a%b
        }
        return a
    }
    lcm := func(a, b int) int {
        return a / gcd(a, b) * b
    }

    st := []int{}
    for _, x := range nums {
        st = append(st, x)
        for len(st) > 1 {
            top := st[len(st)-1]
            prev := st[len(st)-2]
            g := gcd(top, prev)
            if g == 1 {
                break
            }
            st = st[:len(st)-2]
            st = append(st, lcm(top, prev))
        }
    }
    return st   
}

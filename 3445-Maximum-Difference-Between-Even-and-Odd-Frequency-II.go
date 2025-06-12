func maxDifference(s string, k int) int {
    n := len(s)
    res := -1000000
    
    // Try all pairs of characters ('0' to '4')
    for a := '0'; a <= '4'; a++ {
        for b := '0'; b <= '4'; b++ {
            if a == b {
                continue
            }
            
            // Transform string: +1 for char a, -1 for char b, 0 for others
            arr := make([]int, n)
            aa := make([]int, n)  // count of char a
            bb := make([]int, n)  // count of char b
            
            for i := 0; i < n; i++ {
                if rune(s[i]) == a {
                    arr[i] = 1
                    aa[i] = 1
                } else if rune(s[i]) == b {
                    arr[i] = -1
                    bb[i] = 1
                } else {
                    arr[i] = 0
                }
            }
            
            // Build prefix sums
            p := make([]int, n+1)   // prefix sum of differences
            pa := make([]int, n+1)  // prefix sum of char a counts
            pb := make([]int, n+1)  // prefix sum of char b counts
            
            for i := 0; i < n; i++ {
                p[i+1] = p[i] + arr[i]
                pa[i+1] = pa[i] + aa[i]
                pb[i+1] = pb[i] + bb[i]
            }
            
            // Track minimum prefix values for each parity combination
            // mn[i][j] = minimum p[x] where pa[x]%2==i and pb[x]%2==j
            const INF = 1000000
            mn := [2][2]int{{INF, INF}, {INF, INF}}
            
            i := 0
            for j := 0; j <= n; j++ {
                // Sliding window: advance left pointer while valid
                for j-i >= k && pa[j]-pa[i] > 0 && pb[j]-pb[i] > 0 {
                    paIdx := pa[i] % 2
                    pbIdx := pb[i] % 2
                    if p[i] < mn[paIdx][pbIdx] {
                        mn[paIdx][pbIdx] = p[i]
                    }
                    i++
                }

                targetParityA := (pa[j] + 1) % 2  // opposite of current parity
                targetParityB := pb[j] % 2        // same as current parity
                
                if mn[targetParityA][targetParityB] < INF {
                    diff := p[j] - mn[targetParityA][targetParityB]
                    if diff > res {
                        res = diff
                    }
                }
            }
        }
    }
    
    return res
    
}

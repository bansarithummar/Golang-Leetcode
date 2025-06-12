func maxDifference(s string, k int) int {
    n := len(s)
    res := -1000000
    
    for a := '0'; a <= '4'; a++ {
        for b := '0'; b <= '4'; b++ {
            if a == b {
                continue
            }
            
            arr := make([]int, n)
            aa := make([]int, n)  
            bb := make([]int, n)  
            
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
            
            p := make([]int, n+1)  
            pa := make([]int, n+1) 
            pb := make([]int, n+1)  
            
            for i := 0; i < n; i++ {
                p[i+1] = p[i] + arr[i]
                pa[i+1] = pa[i] + aa[i]
                pb[i+1] = pb[i] + bb[i]
            }

            const INF = 1000000
            mn := [2][2]int{{INF, INF}, {INF, INF}}
            
            i := 0
            for j := 0; j <= n; j++ {
                for j-i >= k && pa[j]-pa[i] > 0 && pb[j]-pb[i] > 0 {
                    paIdx := pa[i] % 2
                    pbIdx := pb[i] % 2
                    if p[i] < mn[paIdx][pbIdx] {
                        mn[paIdx][pbIdx] = p[i]
                    }
                    i++
                }

                targetParityA := (pa[j] + 1) % 2 
                targetParityB := pb[j] % 2      
                
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

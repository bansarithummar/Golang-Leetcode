func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
    n := len(startTime)
    gaps := make([]int, n+1)
    gaps[0] = startTime[0]
    for i := 1; i < n; i++ {
        gaps[i] = startTime[i] - endTime[i-1]
    }
    gaps[n] = eventTime - endTime[n-1]
    var v1, v2, v3 int
    var c1, c2, c3 int
    for _, g := range gaps {
        if g > v1 {
            v3, c3 = v2, c2
            v2, c2 = v1, c1
            v1, c1 = g, 1
        } else if g == v1 {
            c1++
        } else if g > v2 {
            v3, c3 = v2, c2
            v2, c2 = g, 1
        } else if g == v2 {
            c2++
        } else if g > v3 {
            v3, c3 = g, 1
        } else if g == v3 {
            c3++
        }
    }
    res := v1
    for i := 0; i < n; i++ {
        a, b := gaps[i], gaps[i+1]
        d := endTime[i] - startTime[i]
        m := a + d + b
        cnt1 := 0
        if a == v1 {
            cnt1++
        }
        if b == v1 {
            cnt1++
        }
        var best int
        if cnt1 < c1 {
            best = v1
        } else {
            cnt2 := 0
            if a == v2 {
                cnt2++
            }
            if b == v2 {
                cnt2++
            }
            if cnt2 < c2 {
                best = v2
            } else {
                best = v3
            }
        }
        var cand int
        if best >= d {
            if m > best {
                cand = m
            } else {
                cand = best
            }
        } else {
            if m-d > best {
                cand = m - d
            } else {
                cand = best
            }
        }
        if cand > res {
            res = cand
        }
    }
    return res
}

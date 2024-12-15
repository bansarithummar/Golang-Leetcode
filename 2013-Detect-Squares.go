type DetectSquares struct {
    count map[[2]int]int
}


func Constructor() DetectSquares {
    return DetectSquares{count: make(map[[2]int]int)} 
}


func (this *DetectSquares) Add(point []int)  {
    p := [2]int{point[0], point[1]}
    this.count[p]++   
}


func (this *DetectSquares) Count(point []int) int {
    x, y := point[0], point[1]
    res := 0
    for p, c := range this.count {
        if p[1] == y && p[0] != x {
            x2 := p[0]
            countSameLine := c
            side := x2 - x

            // Up square
            p1 := [2]int{x, y + side}
            p2 := [2]int{x2, y + side}
            if this.count[p1] > 0 && this.count[p2] > 0 {
                res += countSameLine * this.count[p1] * this.count[p2]
            }

            // Down square (if side != 0, we can also check the other direction)
            p3 := [2]int{x, y - side}
            p4 := [2]int{x2, y - side}
            if this.count[p3] > 0 && this.count[p4] > 0 {
                res += countSameLine * this.count[p3] * this.count[p4]
            }
        }
    }

    return res   
}

func mergeTriplets(triplets [][]int, target []int) bool {
    x, y, z := target[0], target[1], target[2]
    canFormX, canFormY, canFormZ := false, false, false
    for _, triplet := range triplets {
        a, b, c := triplet[0], triplet[1], triplet[2]
        if a <= x && b <= y && c <= z {
            if a == x {
                canFormX = true
            }
            if b == y {
                canFormY = true
            }
            if c == z {
                canFormZ = true
            }
        }
    }
    return canFormX && canFormY && canFormZ   
}

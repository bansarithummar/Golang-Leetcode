217. Contains Duplicate


func containsDuplicate(nums []int) bool {
    hashset := make(map[int]bool)
    for _, n := range nums {
        if hashset[n] {
            return true
        }
        hashset[n] = true
    }
    return false
}


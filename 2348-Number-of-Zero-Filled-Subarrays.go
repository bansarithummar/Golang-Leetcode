func zeroFilledSubarray(nums []int) int64 {
    var res int64 = 0
	var streak int64 = 0

	for _, v := range nums {
		if v == 0 {
			streak++
			res += streak
		} else {
			streak = 0
		}
	}
	return res  
}

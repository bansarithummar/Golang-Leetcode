4. Median of Two Sorted Arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var A, B []int
	A, B = nums1, nums2
	total := len(nums1) + len(nums2)
	half := total / 2

	if len(B) < len(A) {
		A, B = B, A
	}

	l, r := 0, len(A)-1
	for {
		i := (l + r) / 2
		j := half - i - 2

		var Aleft, Aright, Bleft, Bright int
		if i >= 0 {
			Aleft = A[i]
		} else {
			Aleft = math.MinInt64
		}

		if i+1 < len(A) {
			Aright = A[i+1]
		} else {
			Aright = math.MaxInt64
		}

		if j >= 0 {
			Bleft = B[j]
		} else {
			Bleft = math.MinInt64
		}

		if j+1 < len(B) {
			Bright = B[j+1]
		} else {
			Bright = math.MaxInt64
		}

		if Aleft <= Bright && Bleft <= Aright {
			if total%2 == 1 {
				return float64(min(Aright, Bright))
			}
			return float64(max(Aleft, Bleft)+min(Aright, Bright)) / 2
		} else if Aleft > Bright {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

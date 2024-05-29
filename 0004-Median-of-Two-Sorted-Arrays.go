func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	total := m + n
	half := (total + 1) / 2

	l, r := 0, m

	for l <= r {
		i := (l + r) / 2
		j := half - i

		nums1Left := math.Inf(-1)
		if i > 0 {
			nums1Left = float64(nums1[i-1])
		}
		nums1Right := math.Inf(1)
		if i < m {
			nums1Right = float64(nums1[i])
		}

		nums2Left := math.Inf(-1)
		if j > 0 {
			nums2Left = float64(nums2[j-1])
		}
		nums2Right := math.Inf(1)
		if j < n {
			nums2Right = float64(nums2[j])
		}

		if nums1Left <= nums2Right && nums2Left <= nums1Right {
			if total%2 == 0 {
				return (math.Max(nums1Left, nums2Left) + math.Min(nums1Right, nums2Right)) / 2
			} else {
				return math.Max(nums1Left, nums2Left)
			}
		} else if nums1Left > nums2Right {
			r = i - 1
		} else {
			l = i + 1
		}
	}

	panic("Should not reach here")    
}

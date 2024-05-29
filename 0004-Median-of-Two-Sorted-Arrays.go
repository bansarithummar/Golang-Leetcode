func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
    nums1, nums2 = nums2, nums1
    }

    m, n := len(nums1), len(nums2)
    imin, imax, halfLen := 0, m, (m+n+1)/2

    for imin <= imax {
        i := (imin + imax) / 2
        j := halfLen - i

        if i < imax && nums2[j-1] > nums1[i] {
            imin = i + 1
        } else if i > imin && nums1[i-1] > nums2[j] {
            imax = i - 1
        } else {
            // i is perfect
            var maxLeft float64
            if i == 0 {
                maxLeft = float64(nums2[j-1])
            } else if j == 0 {
                maxLeft = float64(nums1[i-1])
            } else {
                maxLeft = float64(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
            }
            if (m+n)%2 == 1 {
                return maxLeft // Odd case
            }

            var minRight float64
            if i == m {
                minRight = float64(nums2[j])
            } else if j == n {
                minRight = float64(nums1[i])
            } else {
                minRight = float64(math.Min(float64(nums1[i]), float64(nums2[j])))
            }

            return (maxLeft + minRight) / 2.0 // Even case
        }
    }

    return 0.0 	
}

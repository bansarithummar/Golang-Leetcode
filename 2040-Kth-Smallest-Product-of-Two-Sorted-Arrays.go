func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
    aSplit := sort.SearchInts(nums1, 0)
	bSplit := sort.SearchInts(nums2, 0)
	a1, a2 := rev(inv(nums1[:aSplit])), nums1[aSplit:]
	b1, b2 := rev(inv(nums2[:bSplit])), nums2[bSplit:]

	negCount := len(a1)*len(b2) + len(b1)*len(a2)

	sign := 1
	if int(k) > negCount {
		k -= int64(negCount) 
	} else {
		k = int64(negCount) - k + 1 
		b1, b2 = b2, b1             
		sign = -1
	}

	var lo, hi int = 0, 1e10
	for lo < hi {
		mid := (lo + hi) / 2
		if count(a1, b1, mid) + count(a2, b2, mid) >= int(k) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return int64(sign * lo)
}

func count(a, b []int, val int) int {
	var count int
	j := len(b) - 1
	for i := range a {
		for j >= 0 && a[i]*b[j] > val {
			j--
		}
		count += j + 1
	}
	return count
}

func inv(a []int) []int {
	for i := range a {
		a[i] = -a[i]
	}
	return a
}

func rev(a []int) []int {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return a    
}

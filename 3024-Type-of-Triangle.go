func triangleType(nums []int) string {
    sort.Ints(nums)

	// Triangle inequality: a + b > c  (after sorting: nums[0] + nums[1] > nums[2])
	if nums[0]+nums[1] <= nums[2] {
		return "none"
	}

	// All three equal → equilateral
	if nums[0] == nums[2] {
		return "equilateral"
	}

	// Exactly two equal → isosceles
	if nums[0] == nums[1] || nums[1] == nums[2] {
		return "isosceles"
	}

	// Otherwise all different → scalene
	return "scalene"
}

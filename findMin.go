153. Find Minimum in Rotated Sorted Array

func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    if len(nums) == 1 || nums[right] > nums[0] {
        return nums[0]
    }
    
    for right >= left {
        mid := left + (right-left)/2
        // If the mid element is greater than its next element then mid+1 element is the smallest
        // This point is the point of change from bigger to smaller
        if nums[mid] > nums[mid+1] {
            return nums[mid+1]
        }
        // If the mid element is lesser than its previous element then mid element is the smallest
        if nums[mid-1] > nums[mid] {
            return nums[mid]
        }
        if nums[mid] > nums[0] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

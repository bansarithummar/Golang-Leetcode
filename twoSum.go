1. Two Sum

package main

func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)
	for i, n := range nums {
		diff := target - n
		if index, found := prevMap[diff]; found {
			return []int{index, i}
		}
		prevMap[n] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // Output: [0, 1]
}

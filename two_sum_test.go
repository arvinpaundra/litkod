package main_test

// https://leetcode.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	dedup := make(map[int]int)

	for currIndex, num := range nums {
		result := target - num

		if prevIndex, exists := dedup[result]; exists {
			return []int{currIndex, prevIndex}
		}

		dedup[result] = currIndex
	}

	return []int{0, 0}
}

package main_test

// https://leetcode.com/problems/contains-duplicate/description/
func containDuplicate(nums []int) bool {
	dedup := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := dedup[num]; exists {
			return true
		}

		dedup[num] = struct{}{}
	}

	return false
}

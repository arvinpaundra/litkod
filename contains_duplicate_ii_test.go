package main_test

// https://leetcode.com/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {
	n := 0
	length := len(nums) - 1

	for n < length {
		if nums[n] == nums[n+1] || abs(n-n+1) <= k {
			return true
		}

		n++
	}

	return false
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

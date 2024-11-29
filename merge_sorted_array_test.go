package main_test

import (
	"sort"
)

// https://leetcode.com/problems/merge-sorted-array/description/
func merge(nums1 []int, _ int, nums2 []int, _ int) {
	l := 0
	r := len(nums1) - len(nums2)

	for r < len(nums1) {
		nums1[r] = nums2[l]

		l++
		r++
	}

	sort.Ints(nums1)
}

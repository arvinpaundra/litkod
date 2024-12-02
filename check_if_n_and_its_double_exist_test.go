package main_test

// https://leetcode.com/problems/check-if-n-and-its-double-exist/description/
func checkIfExist(arr []int) bool {
	hashmap := make(map[int]int)

	n := 0

	for n < len(arr) {
		key := 2 * arr[n]

		if _, ok := hashmap[key]; ok {
			return true
		}

		key = arr[n] / 2

		if _, ok := hashmap[key]; ok && arr[n]%2 == 0 {
			return true
		}

		hashmap[arr[n]] = n

		n++
	}

	return false
}

package main_test

import (
	"slices"
)

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	hashmap := make(map[string][]string)

	for _, str := range strs {
		strBytes := []byte(str)
		slices.Sort(strBytes)
		key := string(strBytes)

		hashmap[key] = append(hashmap[key], str)
	}

	results := make([][]string, len(hashmap))
	index := 0

	for _, anagrams := range hashmap {
		results[index] = anagrams
		index++
	}

	return results
}

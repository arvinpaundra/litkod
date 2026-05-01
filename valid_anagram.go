package main_test

// https://leetcode.com/problems/valid-anagram/
func validAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap, tMap := make(map[rune]int16), make(map[rune]int16)

	for _, val := range s {
		sMap[val]++
	}

	for _, val := range t {
		tMap[val]++
	}

	if len(sMap) != len(tMap) {
		return false
	}

	for key, val := range sMap {
		if val != tMap[key] {
			return false
		}
	}

	return true
}

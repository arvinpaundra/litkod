package main_test

// https://leetcode.com/problems/detect-capital/description/
func DetectCapitalUse(word string) bool {
	n := 0
	capital := 0
	lower := 0

	for n < len(word) {
		if word[n] >= 'A' && word[n] <= 'Z' {
			capital++
		}

		if word[n] >= 'a' && word[n] <= 'z' {
			lower++
		}

		n++
	}

	if len(word) == capital {
		return true
	}

	if len(word) == lower {
		return true
	}

	if word[0] >= 'A' && word[0] <= 'Z' && capital == 1 {
		return true
	}

	return false
}

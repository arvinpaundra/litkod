package main_test

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func StrStr(haystack, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	lNeedle := 0
	rNeedle := len(needle)

	for rNeedle <= len(haystack) {
		currentHaystack := haystack[lNeedle:rNeedle]

		if currentHaystack == needle {
			return lNeedle
		}

		lNeedle++
		rNeedle++
	}

	return -1
}

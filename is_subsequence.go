package main_test

// https://leetcode.com/problems/is-subsequence/description/
func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(t) {
		if j < len(s) && t[i] == s[j] {
			j++
		}

		i++
	}

	return j == len(s)
}

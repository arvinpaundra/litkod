package main_test

// https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/
func canMakeSubsequence(str1 string, str2 string) bool {
	i, j := 0, 0

	for i < len(str1) && j < len(str2) {
		char := str1[i] + 1

		if char > 'z' {
			char = 'a'
		}

		if str1[i] == str2[j] || char == str2[j] {
			j++
		}

		i++
	}

	return j == len(str2)
}

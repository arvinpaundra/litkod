package main_test

import (
	"strings"
)

// https://leetcode.com/problems/valid-palindrome/description/
func isPalindrome(s string) bool {
	return sanitize(s) == reverse(sanitize(s))
}

func sanitize(s string) string {
	var b []byte

	n := 0

	for n < len(s) {
		if s[n] >= 'A' && s[n] <= 'Z' {
			b = append(b, s[n])
		} else if s[n] >= 'a' && s[n] <= 'z' {
			b = append(b, s[n])
		} else if s[n] >= '0' && s[n] <= '9' {
			b = append(b, s[n])
		}

		n++
	}

	return strings.ToLower(string(b))
}

func reverse(str string) string {
	n := len(str) - 1

	var result []byte

	for n >= 0 {
		result = append(result, str[n])

		n--
	}

	return string(result)
}

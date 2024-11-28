package main_test

var romans = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// https://leetcode.com/problems/roman-to-integer/description/
func romanToInt(s string) int {
	result, n := 0, len(s)-1

	for n >= 0 {
		key := s[n]

		if n > 0 {
			seek := s[n-1]

			if romans[seek] < romans[key] {
				result += romans[key] - romans[seek]

				n -= 2

				continue
			}
		}

		result += romans[key]

		n--
	}

	return result
}

package main_test

// https://leetcode.com/problems/adding-spaces-to-a-string/
func addSpaces(s string, spaces []int) string {
	i, j := 0, 0
	b := make([]byte, 0, len(s)+len(spaces))

	for i < len(s) {
		if j < len(spaces) && i == spaces[j] {
			b = append(b, ' ')

			j++
		}

		b = append(b, s[i])

		i++
	}

	return string(b)
}

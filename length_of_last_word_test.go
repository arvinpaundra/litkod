package main_test

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{
			str:      "Hello World",
			expected: 5,
		},
		{
			str:      "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			str:      "luffy is still joyboy",
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			res := LengthOfLastWord(test.str)

			if res != test.expected {
				t.Errorf("expected %d, but got %d", test.expected, res)
				t.FailNow()
			}
		})
	}
}

func LengthOfLastWord(s string) int {
	result := 0
	n := len(s)

	for n > 0 {
		n--

		if s[n] != ' ' {
			result++
		} else if result > 0 {
			return result
		}
	}

	return result
}

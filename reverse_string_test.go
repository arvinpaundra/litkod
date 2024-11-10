package main_test

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		in       []byte
		expected string
	}{
		{
			in:       []byte{'h', 'e', 'l', 'l', 'o'},
			expected: "olleh",
		},
		{
			in:       []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expected: "hannaH",
		},
	}

	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			ReverseString(test.in)

			if string(test.in) != test.expected {
				t.Errorf("expected %s, but got %s", test.expected, string(test.in))
				t.FailNow()
			}
		})
	}
}

func ReverseString(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

package main_test

import (
	"strings"
	"testing"
)

func reverseWordsInAString(s string) string {
	trim := strings.TrimSpace(s)
	words := strings.Fields(trim)

	start, end := 0, len(words)-1

	for start < end {
		words[start], words[end] = words[end], words[start]

		start++
		end--
	}

	return strings.Join(words, " ")
}

func TestReverseWordsInAString(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{
			in:       "the sky is blue",
			expected: "blue is sky the",
		},
		{
			in:       "  hello world  ",
			expected: "world hello",
		},
		{
			in:       "a good   example",
			expected: "example good a",
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			result := reverseWordsInAString(tt.in)

			if result != tt.expected {
				t.Fatalf("expected %s, but got %s", tt.expected, result)
			}
		})
	}
}

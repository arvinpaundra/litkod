package main_test

import "testing"

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected bool
	}{
		{
			name:     "success",
			str:      "()",
			expected: true,
		},
		{
			name:     "success",
			str:      "()[]{}",
			expected: true,
		},
		{
			name:     "failed",
			str:      "(]",
			expected: false,
		},
		{
			name:     "success",
			str:      "([])",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := IsValidParentheses(test.str)

			if res != test.expected {
				t.Errorf("expected %v, but got %v", test.expected, res)
				t.FailNow()
			}
		})
	}
}

func IsValidParentheses(s string) bool {
	parentheses := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	var stack []byte

	for _, str := range s {
		switch str {
		case '{', '(', '[':
			stack = append(stack, byte(str))
		case '}', ')', ']':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if parentheses[byte(str)] != top {
				return false
			}

			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

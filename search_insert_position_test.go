package main_test

import (
	"testing"
)

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		in       []int
		target   int
		expected int
	}{
		{
			in:       []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			in:       []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			in:       []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			res := SearchInsert(test.in, test.target)

			if res != test.expected {
				t.Errorf("expected %d, but got %d", test.expected, res)
				t.FailNow()
			}
		})
	}
}

func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		v := nums[mid]

		switch {
		case v < target:
			left = mid + 1
		case v > target:
			right = mid - 1
		default:
			return mid
		}
	}

	return left
}

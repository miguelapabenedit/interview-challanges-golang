package challanges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsoluteDifferenceBetweenDiagonals_EmptySquare_ReturnsError(t *testing.T) {
	var square [][]int
	var expected = "the square cant be of lenght 0"
	_, e := AbsoluteDifferenceBetweenDiagonals(square)

	assert.True(t, e.Error() == expected)
}

func TestAbsoluteDifferenceBetweenDiagonals_InvalidSquareLenght_ReturnsError(t *testing.T) {
	var square [][]int = [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}

	var expected = "the multidemnsional array provided is not a square"
	_, e := AbsoluteDifferenceBetweenDiagonals(square)

	assert.True(t, e.Error() == expected)
}

func TestAbsoluteDifferenceBetweenDiagonals_ValidSquare_ReturnsAbsoluteInt(t *testing.T) {
	type SquareResult struct {
		matrix   [][]int
		expected int
	}

	testCases := []SquareResult{
		{
			matrix: [][]int{
				{1, 2},
				{1, 2},
			},
			expected: 1,
		}, {
			matrix: [][]int{
				{1, 3, 3, 3},
				{1, 21, 23, 2},
				{1, 21, 1, 2},
				{1, 21, 1, 2},
			},
			expected: 16,
		},
	}

	for _, val := range testCases {
		r, _ := AbsoluteDifferenceBetweenDiagonals(val.matrix)
		assert.Equal(t, val.expected, r)
	}
}

func TestIsPalindrome(t *testing.T) {
	type TestCases struct {
		s        string
		expected bool
	}

	testCases := []TestCases{
		{
			"menem",
			true,
		},
		{
			"",
			true,
		},
		{
			"pamela",
			false,
		},
	}

	for _, v := range testCases {
		r := IsPalindrome(v.s)
		assert.Equal(t, v.expected, r)
	}
}

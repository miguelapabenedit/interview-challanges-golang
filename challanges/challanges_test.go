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
		matrix   [2][2]int
		expected int
	}

	testCases := []SquareResult{
		{
			matrix: [2][2]int{
				{1, 2},
				{1, 2},
			},
			expected: 0,
		}, {
			matrix: [2][2]int{
				{1, 2},
				{1, 2},
			},
			expected: 0,
		},
	}

	for _, val := range testCases {
		r, _ := AbsoluteDifferenceBetweenDiagonals(val.matrix)
		assert.True(t, r == val.expected)
	}
}

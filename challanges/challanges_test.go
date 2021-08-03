package challanges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeConvertion(t *testing.T) {
	type TestCase struct {
		d string
		e string
	}

	tc := []TestCase{
		{"11:00:00AM", "11:00:00"},
		{"01:00:00AM", "01:00:00"},
		{"02:00:00AM", "02:00:00"},
		{"03:00:00AM", "03:00:00"},
		{"04:00:00AM", "04:00:00"},
		{"05:00:00AM", "05:00:00"},
		{"06:00:00AM", "06:00:00"},
		{"07:00:00AM", "07:00:00"},
		{"08:00:00AM", "08:00:00"},
		{"09:00:00AM", "09:00:00"},
		{"10:00:00AM", "10:00:00"},
		{"12:00:00AM", "00:00:00"},
		{"01:00:00PM", "13:00:00"},
		{"02:00:00PM", "14:00:00"},
		{"03:00:00PM", "15:00:00"},
		{"04:00:00PM", "16:00:00"},
		{"05:00:00PM", "17:00:00"},
		{"06:00:00PM", "18:00:00"},
		{"07:00:00PM", "19:00:00"},
		{"08:00:00PM", "20:00:00"},
		{"09:00:00PM", "21:00:00"},
		{"10:00:00PM", "22:00:00"},
		{"11:00:00PM", "23:00:00"},
		{"12:00:00PM", "12:00:00"},
	}

	for _, val := range tc {
		r := TimeConversion(val.d)
		assert.Equal(t, val.e, r)
	}
}

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

func TestCalculateBinaryGap(t *testing.T) {
	type TestCases struct {
		s        int
		expected int
	}

	testCases := []TestCases{
		{
			512487,
			3,
		},
		{
			203,
			2,
		},
		{
			21,
			1,
		},
	}

	for _, v := range testCases {
		r := CalculateBinaryGap(v.s)
		assert.Equal(t, v.expected, r)
	}
}

func TestRotateRightArray(t *testing.T) {
	type TestCases struct {
		s        []int
		n        int
		expected []int
	}

	testCases := []TestCases{
		{
			[]int{1, 2, 3, 4, 5, 6},
			6,
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			3,
			[]int{4, 5, 6, 1, 2, 3},
		},
		{
			[]int{},
			4,
			[]int{},
		},
	}

	for _, v := range testCases {
		r := RotateRightArray(v.s, v.n)
		assert.Equal(t, v.expected, r)
	}
}

func TestFindFirstOddOcurrenceInArray(t *testing.T) {
	type TestCases struct {
		s        []int
		expected int
	}

	testCases := []TestCases{
		{
			[]int{1, 1, 3, 3, 5, 5, 4},
			4,
		},
		{
			[]int{1, 1, 3, 3, 5, 5, 4, 4, 9},
			9,
		},
		{
			[]int{4},
			4,
		},
	}

	for _, v := range testCases {
		r := FindFirstOddOcurrenceInArray(v.s)
		assert.Equal(t, v.expected, r)
	}
}

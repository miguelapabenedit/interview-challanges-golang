package challanges

import (
	"errors"
	"math"
)

// AbsoluteDifferenceBetweenDiagonals calculates the absolute difference between two diagonals of a square matrix.
// It returns int as the absolute difference
func AbsoluteDifferenceBetweenDiagonals(square [][]int) (int, error) {
	if len(square) == 0 {
		return 0, errors.New("the square cant be of lenght 0")
	}

	if len(square) != len(square[0]) {
		return 0, errors.New("the multidemnsional array provided is not a square")
	}

	d1, d2 := 0, 0

	for i := 0; i < len(square); i++ {
		d1 += square[i][i]
		d2 += square[i][len(square)-1]
	}

	sr := d1 - d2
	r := math.Abs(float64(sr))

	return int(r), nil
}

func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

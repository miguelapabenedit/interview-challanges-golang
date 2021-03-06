package challanges

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//https://www.hackerrank.com/challenges/apple-and-orange/problem
func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	orangeCount, appleCount := 0, 0

	for i := 0; i < len(apples); i++ {
		al := a + apples[i]
		if al >= s && al <= t {
			appleCount++
		}
	}

	for i := 0; i < len(oranges); i++ {
		or := b + oranges[i]
		if or <= t && or >= s {
			orangeCount++
		}
	}

	fmt.Println(appleCount)
	fmt.Println(orangeCount)
}

func GradientStudent(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		if grades[i] > 37 {
			r := (((grades[i] / 5) + 1) * 5)
			if (r - grades[i]) < 3 {
				grades[i] = r
			}
		}
	}
	return grades
}

func TimeConversion(date string) string {
	date, zone := date[:len(date)-2], date[len(date)-2:]
	h, _ := strconv.Atoi(date[:2])
	finalH := date[:2]

	if strings.ToUpper(zone) == "PM" {
		if h != 12 {
			finalH = strconv.Itoa(h + 12)
		}
	} else if h == 12 {
		finalH = "00"
	}

	return fmt.Sprintf("%v%v", finalH, date[2:])
}

func BirthDayCakeCandles(arr []int) {
	var min int64
	var max int64
	var sum int64

	for i := 0; i < len(arr); i++ {
		sum = sum + int64(arr[i])

		if int64(arr[i]) < min || min == 0 {
			min = int64(arr[i])
		}
		if int64(arr[i]) > max {
			max = int64(arr[i])
		}
	}

	fmt.Printf("%v %v", sum-max, sum-min)
}

func Staircase(n int32) {
	for i := 1; i <= int(n); i++ {
		fmt.Print(strings.Repeat(" ", int(n)-i), strings.Repeat("#", i), "\n")
	}
}

func PlusMinus() {
	var negatives = 1
	var positives = 1
	var ceroes = 1

	arr := []int{1, 0, -1}
	for _, v := range arr {
		switch {
		case v > 0:
			positives += 1
		case v < 0:
			negatives += 1
		default:
			ceroes += 1
		}
	}

	fmt.Println(float32(positives) / float32(len(arr)))
	fmt.Println(float32(negatives) / float32(len(arr)))
	fmt.Println(float32(ceroes) / float32(len(arr)))

}

// Codility: finds the first unpaired number (same value)
// Returns odd int
func FindFirstOddOcurrenceInArray(A []int) int {
	pairs := make(map[int]bool)
	result := 0

	for i := 0; i < len(A); i++ {
		_, ok := pairs[A[i]]
		if ok {
			delete(pairs, A[i])
		} else {
			pairs[A[i]] = true
		}
	}

	for _, v := range A {
		if pairs[v] {
			result = v
			break
		}
	}

	return result
}

// Codility: rotates array right N times
// Returns the rotated array
func RotateRightArray(A []int, K int) []int {

	if len(A) > 1 {
		for i := 0; i < K; i++ {
			A = append([]int{A[len(A)-1]}, A[:len(A)-1]...)
		}
	}

	return A
}

// Codility: A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.
// It returns int as the maximal sequence

func CalculateBinaryGap(N int) int {
	acumulator, maxSecuence := 0, 0
	firstCheckpoint := false

	for _, code := range strconv.FormatInt(int64(N), 2) {
		if code == 49 {
			firstCheckpoint = true
			if firstCheckpoint {
				if acumulator > maxSecuence {
					maxSecuence = acumulator
				}
				acumulator = 0
			}

		} else {
			acumulator++
		}
	}

	return maxSecuence
}

// Hackerank: AbsoluteDifferenceBetweenDiagonals calculates the absolute difference between two diagonals of a square matrix.
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

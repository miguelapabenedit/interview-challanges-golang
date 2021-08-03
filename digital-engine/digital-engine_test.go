package digitalengine_test

import (
	"fmt"
	"testing"

	digitalengine "github.com/miguelapabenedit/interview-challanges-golang/digital-engine"
)

func TestSolu2(t *testing.T) {
	type TestCase struct {
		i []int
		e int
	}
	tc := []TestCase{
		{[]int{1, 3, 8, 3, 1}, 3},
		{[]int{1, 3, 8, 3, 1}, 3},
		{[]int{7, 3, 7, 3, 1, 3, 4, 1}, 5},
		{[]int{7, 2, 2, 9, 9, 8, 3, 2, 2, 3, 1, 7}, 8},
		{[]int{1, 2, 3, 4, 1}, 4},
		{[]int{1, 3, 8, 3, 4, 2}, 6},
		{[]int{7, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3}, 24},
		{[]int{0}, 1},
	}

	for _, v := range tc {
		got := digitalengine.TravelerProblem(v.i)
		fmt.Println(got)
		if got != v.e {
			t.Errorf("got :%v expected:%v", got, v.e)
		}
	}
}

func TestSolut1(t *testing.T) {
	// type testCase struct {
	// 	i int
	// 	e int
	// }

	// tc := []testCase{
	// 	{12345, 54321},
	// 	{00001, 1},
	// 	{10000, 1},
	// 	{12, 21},
	// 	{20001, 10002},
	// 	{100233021, 120332001},
	// }

	// for _, v := range tc {
	// 	//	got := digitalengine.Solve1(v.i)
	// 	if got != v.e {
	// 		t.Errorf("got :%v expected:%v", got, v.e)
	// 	}
	// }
}

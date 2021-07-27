package digitalengine

import "fmt"

//reverse a number 54321 12345 400 4 0004 4
func Problem1(N int) {
	var enable_print int
	enable_print = N % 10
	for N > 0 {
		if enable_print == 0 && N%10 != 0 {
			enable_print = 1
		} else if enable_print == 1 {
			fmt.Print(N % 10)
		}

		N = N / 10
	}
}

package digitalengine

import "fmt"

func ReverseNumberByDivide(N int) {
	var enable_print int
	enable_print = N % 10
	for N > 0 {
		if enable_print == 0 && N%10 != 0 {
			enable_print = 1
		}
		if enable_print != 0 {
			fmt.Print(N % 10)
		}

		N = N / 10
	}
}

func TravelerProblem(arr []int) int {
	destinations := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if _, ok := destinations[arr[i]]; !ok {
			destinations[arr[i]] = i
		}
	}

	steps := len(arr)
	for i := 0; i < len(arr) && len(destinations) <= len(arr)-i; i++ {
		visits := make(map[int]int)
		counterSteps, unsavedSteps := 0, 0

		for j := i; j < len(arr); j++ {
			unsavedSteps++
			if _, ok := visits[arr[j]]; !ok {
				visits[arr[j]] = 1
				counterSteps = counterSteps + unsavedSteps
				unsavedSteps = 0
			}
		}

		hasTravelAll := true
		for k, _ := range destinations {
			if _, ok := visits[k]; !ok {
				hasTravelAll = false
				break
			}
		}

		if counterSteps < steps && hasTravelAll {
			steps = counterSteps
		}
	}

	return steps
}

func TravelerProblem2(arr []int) int {
	destinations := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		destinations[arr[i]] += 1
	}

	start := -1
	end := 0

	for i := 0; i < len(arr) && start < 0; i++ {
		visits := make(map[int]int)

		if r := destinations[arr[i]]; r == 1 {
			start = i

			for j := i; j < len(arr); j++ {
				if _, ok := visits[arr[j]]; !ok {
					visits[arr[j]] = 1
					end = j + 1
				}
			}
		} else {
			destinations[arr[i]] = destinations[arr[i]] - 1
		}
	}

	return (end - start)
}

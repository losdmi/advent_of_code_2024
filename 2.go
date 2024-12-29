package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func secondMain() error {
	rawInput, err := os.ReadFile("2.txt")
	if err != nil {
		return err
	}
	rawLines := strings.Split(strings.TrimSpace(string(rawInput)), "\n")

	reports := make([][]int, len(rawLines))
	for i, rawLine := range rawLines {
		fields := strings.Fields(rawLine)

		report := make([]int, len(fields))
		for i := range fields {
			report[i], _ = strconv.Atoi(fields[i])
		}
		reports[i] = report
	}

	var result int
	// result = secondFirstSolution(reports)
	// fmt.Println(result)
	result = secondSecondSolution(reports)
	fmt.Println(result)

	return nil
}

func secondFirstSolution(reports [][]int) int {
	var result int
	for i, report := range reports {
		isSafe := isSafeReport(report)
		if isSafe {
			result += 1
			fmt.Println(i+1, "Safe")
		} else {
			fmt.Println(i+1, "Unsafe")
		}
	}

	return result
}

func isSafeReport(report []int) bool {
	if len(report) == 1 {
		return true
	}
	if len(report) == 2 {
		return math.Abs(float64(report[0]-report[1])) < 4
	}

	for i := 2; i < len(report); i++ {
		prevDiff := report[i-1] - report[i-2]
		curDiff := report[i] - report[i-1]
		if prevDiff*curDiff <= 0 || math.Abs(float64(prevDiff)) > 3 || math.Abs(float64(curDiff)) > 3 {
			return false
		}
	}

	return true
}

func secondSecondSolution(reports [][]int) int {
	var result int
	for i, report := range reports {
		// fmt.Println()
		isSafe := isSafeReport(report)
		if isSafe {
			result += 1
			fmt.Println(i+1, "Safe")
		} else {
			isSafe = false
			for j := 0; j < len(report); j++ {
				withoutOne := make([]int, 0, len(report)-1)
				withoutOne = append(withoutOne, report[:j]...)
				withoutOne = append(withoutOne, report[j+1:]...)

				// fmt.Println(withoutOne)

				isSafe = isSafeReport(withoutOne)
				if isSafe {
					break
				}
			}
			if isSafe {
				result += 1
				fmt.Println(i+1, "Safe")
			} else {
				fmt.Println(i+1, "Unsafe")
			}
		}
	}

	return result
}

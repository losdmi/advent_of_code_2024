package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func firstFirstMain() error {
	rawInput, err := os.ReadFile("1-1.txt")
	if err != nil {
		return err
	}
	rawLines := strings.Split(strings.TrimSpace(string(rawInput)), "\n")

	listA, listB := make([]int, len(rawLines)), make([]int, len(rawLines))
	for i, rawLine := range rawLines {
		fields := strings.Fields(rawLine)
		listA[i], err = strconv.Atoi(fields[0])
		if err != nil {
			return err
		}
		listB[i], err = strconv.Atoi(fields[1])
		if err != nil {
			return err
		}
	}

	// result := firstFirstSolution(listA, listB)
	result := firstSecondSolution(listA, listB)
	fmt.Println(result)

	return nil
}

func firstFirstSolution(listA, listB []int) int {
	slices.Sort(listA)
	slices.Sort(listB)

	var result int
	for i := 0; i < len(listA); i++ {
		result += int(math.Abs(float64(listA[i] - listB[i])))
	}

	return result
}

func firstSecondSolution(listA, listB []int) int {
	m := make(map[int]int)

	for i := 0; i < len(listB); i++ {
		m[listB[i]] += 1
	}

	var result int
	for _, n := range listA {
		result += n * m[n]
	}

	return result
}

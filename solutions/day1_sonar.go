package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"strconv"
)

func Day1() error {
	var input []int

	err := helpers.ReadLines("inputs/day1.txt", func(s string) error {
		depth, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("input file parse err: %w", err)
		}

		input = append(input, depth)
		return nil
	})
	if err != nil {
		return err
	}

	depthIncreases := NumberOfIncreasesInIntSlice(input)
	fmt.Printf("Day 1 Part 1: %d\n", depthIncreases)

	depthIncreases = SlidingNumberOfIncreasesInIntSlice(input)
	fmt.Printf("Day 1 Part 2: %d\n", depthIncreases)

	return nil
}

func NumberOfIncreasesInIntSlice(input []int) int {
	lastDepth := input[0]
	increased := 0
	for i := 1; i < len(input); i++ {
		if lastDepth < input[i] {
			increased++
		}

		lastDepth = input[i]
	}
	return increased
}

func SlidingNumberOfIncreasesInIntSlice(input []int) int {
	if len(input) < 4 {
		return 0
	}

	increased := 0
	for i := 3; i < len(input); i++ {
		if helpers.SliceSum(input[i-3:i]) < helpers.SliceSum(input[i-2:i+1]) {
			increased++
		}
	}

	return increased
}

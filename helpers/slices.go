package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func StringsToInts(strs []string) []int {
	ints := make([]int, 0, len(strs))

	for _, str := range strs {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		ints = append(ints, i)
	}
	return ints
}

func IntsMinMax(slice []int) (int, int) {
	min, max := slice[0], slice[0]

	for _, i := range slice {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	return min, max
}

func MulIntSlice(slice []int) int {
	mul := 1

	for _, v := range slice {
		mul *= v
	}

	return mul
}

func SliceSum(slice []int) int {
	collector := 0
	for _, v := range slice {
		collector += v
	}

	return collector
}

func IndexOfInt(slice []int, value int) (int, bool) {
	for i, v := range slice {
		if v == value {
			return i, true
		}
	}

	return -1, false
}

func IndexOfString(slice []string, value string) (int, bool) {
	for i, v := range slice {
		if v == value {
			return i, true
		}
	}

	return -1, false
}

func StringSliceContains(slice []string, value string) bool {
	_, ok := IndexOfString(slice, value)
	return ok
}

func FilterEmptyStrings(strs []string) []string {
	filtered := make([]string, 0, len(strs))

	for _, str := range strs {
		if strings.TrimSpace(str) != "" {
			filtered = append(filtered, str)
		}
	}

	return filtered
}

func PrintIntGrid(grid [][]int) {
	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			fmt.Printf("%d", grid[x][y])
		}
		fmt.Println()
	}
}

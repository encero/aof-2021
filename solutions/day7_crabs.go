package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"strings"
)

func Day7Crabs() error {
	input := helpers.StringsToInts(
		strings.Split(
			helpers.ReadLine("inputs/day7.txt"),
			","),
	)

	optimum := CrabPosition(input, func(i int) int {
		return i
	})

	fmt.Println("Day 7 - Part 1:", optimum)

	optimum = CrabPosition(input, CrabDistance)

	fmt.Println("Day 7 - Part 2:", optimum)

	return nil
}

func SliceDistance(target int, slice []int, distFunc func(int) int) int {
	var distance int

	for _, v := range slice {
		distance += distFunc(helpers.AbsInt(v - target))
	}

	return distance
}

func CrabPosition(slice []int, distFunc func(int) int) int {
	min, max := helpers.IntsMinMax(slice)

	var minDistance = SliceDistance(min, slice, distFunc)
	for i := min; i <= max; i++ {
		distance := SliceDistance(i, slice, distFunc)
		if distance < minDistance {
			minDistance = distance
		}
	}

	return minDistance
}

func CrabDistance(i int) int {
	return helpers.SumOfSeries(1, i+1)
}

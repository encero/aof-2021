package solutions

import (
	"github.com/encero/advent-of-code-2021/helpers"
	is_ "github.com/matryer/is"
	"sort"
	"testing"
)

var smokeTestInput = [][]int{
	{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
	{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
	{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
	{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
}

func TestLowestSmokePoints(t *testing.T) {
	is := is_.New(t)

	points := LowestSmokePoints(smokeTestInput)

	var heights []int
	for _, point := range points {
		heights = append(heights, smokeTestInput[point.X][point.Y])
	}

	sort.Ints(heights)

	is.Equal(heights, []int{0, 1, 5, 5})
}

func TestCalculateSmokeBasin(t *testing.T) {
	is := is_.New(t)

	var sizes []int

	points := LowestSmokePoints(smokeTestInput)

	for _, p := range points {
		sizes = append(sizes, CalculateSmokeBasin(p, smokeTestInput))
	}

	sort.Ints(sizes)

	subSlice := sizes[len(sizes)-3:]

	result := helpers.MulIntSlice(subSlice)

	is.Equal(result, 1134)
}

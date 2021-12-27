package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"sort"
	"strings"
)

func Day9Smoke() error {
	var input [][]int
	helpers.ReadLines("inputs/day9.txt", func(s string) error {
		row := helpers.StringsToInts(strings.Split(s, ""))
		input = append(input, row)

		return nil
	})

	lowPoints := LowestSmokePoints(input)

	var riskLevel int
	for _, point := range lowPoints {
		riskLevel += input[point.X][point.Y]
	}

	fmt.Println("Day 9 part 1:", riskLevel)

	var sizes []int

	for _, p := range lowPoints {
		sizes = append(sizes, CalculateSmokeBasin(p, input))
	}

	sort.Ints(sizes)

	subSlice := sizes[len(sizes)-3:]

	result := helpers.MulIntSlice(subSlice)

	fmt.Println("Day 9 part 2:", result)

	return nil
}

func ArrayNeighbors(input [][]int, x, y int) []int {
	nei := make([]int, 0, 4)

	if x > 0 {
		nei = append(nei, input[x-1][y])
	}

	if y > 0 {
		nei = append(nei, input[x][y-1])
	}

	if x < len(input)-1 {
		nei = append(nei, input[x+1][y])
	}

	if y < len(input[0])-1 {
		nei = append(nei, input[x][y+1])
	}

	return nei
}

func LowestSmokePoints(input [][]int) []Point {
	var out []Point

	for x := 0; x < len(input); x++ {

	loop:
		for y := 0; y < len(input[x]); y++ {
			value := input[x][y]

			nei := ArrayNeighbors(input, x, y)
			for _, n := range nei {
				if n <= value {
					continue loop
				}
			}

			out = append(out, Point{x, y})
		}
	}

	return out
}

type Point struct {
	X int
	Y int
}

func BiggerNeighbors(p Point, input [][]int) []Point {
	value := input[p.X][p.Y]

	var points []Point

	var x, y int

	x, y = p.X-1, p.Y
	if p.X > 0 && input[x][y] > value && input[x][y] < 9 {
		points = append(points, Point{x, y})
	}

	x, y = p.X, p.Y-1
	if p.Y > 0 && input[x][y] > value && input[x][y] < 9 {
		points = append(points, Point{x, y})
	}

	x, y = p.X+1, p.Y
	if p.X < len(input)-1 && input[x][y] > value && input[x][y] < 9 {
		points = append(points, Point{x, y})
	}

	x, y = p.X, p.Y+1
	if p.Y < len(input[0])-1 && input[x][y] > value && input[x][y] < 9 {
		points = append(points, Point{x, y})
	}

	return points
}

func CalculateSmokeBasin(start Point, input [][]int) int {
	visited := make(map[Point]struct{})

	visit := []Point{start}

	for len(visit) > 0 {
		point := visit[0]
		visit = visit[1:]

		if _, ok := visited[point]; ok {
			continue
		}

		visited[point] = struct{}{}

		nei := BiggerNeighbors(point, input)

		visit = append(visit, nei...)
	}

	return len(visited)
}

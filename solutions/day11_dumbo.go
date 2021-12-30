package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
)

func Day11DumboOcto() error {
	var input = [][]DumboOctopus{
		{{PowerLevel: 6}, {PowerLevel: 6}, {PowerLevel: 1}, {PowerLevel: 7}, {PowerLevel: 1}, {PowerLevel: 1}, {PowerLevel: 3}, {PowerLevel: 5}, {PowerLevel: 8}, {PowerLevel: 4}},
		{{PowerLevel: 6}, {PowerLevel: 5}, {PowerLevel: 4}, {PowerLevel: 4}, {PowerLevel: 2}, {PowerLevel: 1}, {PowerLevel: 8}, {PowerLevel: 6}, {PowerLevel: 3}, {PowerLevel: 8}},
		{{PowerLevel: 5}, {PowerLevel: 4}, {PowerLevel: 5}, {PowerLevel: 7}, {PowerLevel: 3}, {PowerLevel: 3}, {PowerLevel: 1}, {PowerLevel: 4}, {PowerLevel: 8}, {PowerLevel: 8}},
		{{PowerLevel: 1}, {PowerLevel: 1}, {PowerLevel: 3}, {PowerLevel: 5}, {PowerLevel: 6}, {PowerLevel: 7}, {PowerLevel: 5}, {PowerLevel: 5}, {PowerLevel: 8}, {PowerLevel: 7}},
		{{PowerLevel: 1}, {PowerLevel: 2}, {PowerLevel: 2}, {PowerLevel: 1}, {PowerLevel: 3}, {PowerLevel: 5}, {PowerLevel: 3}, {PowerLevel: 2}, {PowerLevel: 1}, {PowerLevel: 6}},
		{{PowerLevel: 1}, {PowerLevel: 8}, {PowerLevel: 1}, {PowerLevel: 1}, {PowerLevel: 1}, {PowerLevel: 2}, {PowerLevel: 4}, {PowerLevel: 3}, {PowerLevel: 7}, {PowerLevel: 8}},
		{{PowerLevel: 1}, {PowerLevel: 3}, {PowerLevel: 8}, {PowerLevel: 7}, {PowerLevel: 8}, {PowerLevel: 6}, {PowerLevel: 4}, {PowerLevel: 3}, {PowerLevel: 6}, {PowerLevel: 8}},
		{{PowerLevel: 4}, {PowerLevel: 4}, {PowerLevel: 2}, {PowerLevel: 7}, {PowerLevel: 6}, {PowerLevel: 3}, {PowerLevel: 7}, {PowerLevel: 2}, {PowerLevel: 6}, {PowerLevel: 2}},
		{{PowerLevel: 6}, {PowerLevel: 7}, {PowerLevel: 7}, {PowerLevel: 8}, {PowerLevel: 6}, {PowerLevel: 4}, {PowerLevel: 5}, {PowerLevel: 4}, {PowerLevel: 8}, {PowerLevel: 6}},
		{{PowerLevel: 3}, {PowerLevel: 6}, {PowerLevel: 8}, {PowerLevel: 2}, {PowerLevel: 1}, {PowerLevel: 4}, {PowerLevel: 6}, {PowerLevel: 7}, {PowerLevel: 4}, {PowerLevel: 5}},
	}

	simulator := NewDumboSimulator(input)

	for count := 0; count < 100; count++ {
		simulator.Step()
	}

	fmt.Println("Day 11 Part1: ", simulator.Flashes())

	for {
		simulator.Step()
		if simulator.LastStepFlashes() == 100 {
			break
		}
	}

	fmt.Println("Day 11 Part 2: ", simulator.StepCount())

	return nil
}

type DumboOctopus struct {
	PowerLevel      int
	RecentlyFlashed bool
}

type DumboSimulator struct {
	grid            [][]DumboOctopus
	flashCount      int
	simulationSteps int
	lastFlashCount  int
}

func NewDumboSimulator(grid [][]DumboOctopus) *DumboSimulator {
	return &DumboSimulator{
		grid:       grid,
		flashCount: 0,
	}
}

func (d *DumboSimulator) Step() {
	d.lastFlashCount = 0
	d.simulationSteps++

	flashy := d.raisePowerLevels()

	if len(flashy) == 0 {
		return
	}

	d.flashTheOctos(flashy)
}

func (d *DumboSimulator) Flashes() int {
	return d.flashCount
}

func (d DumboSimulator) LastStepFlashes() int {
	return d.lastFlashCount
}

func (d DumboSimulator) StepCount() int {
	return d.simulationSteps
}

func (d *DumboSimulator) raisePowerLevels() []Point {
	var flashy []Point

	for x := 0; x < len(d.grid); x++ {
		for y := 0; y < len(d.grid[x]); y++ {
			d.grid[x][y].PowerLevel += 1
			d.grid[x][y].RecentlyFlashed = false

			if d.grid[x][y].PowerLevel > 9 {
				flashy = append(flashy, Point{x, y})
			}
		}
	}

	return flashy
}

func (d *DumboSimulator) flashTheOctos(flashy []Point) {
	for len(flashy) > 0 {
		nei := d.flashOne(flashy[0])

		flashy = append(flashy, nei...)

		flashy = flashy[1:]
	}
}

func (d *DumboSimulator) flashOne(point Point) []Point {
	theOne := d.grid[point.X][point.Y]
	if theOne.RecentlyFlashed {
		return []Point{}
	}

	if theOne.PowerLevel <= 9 {
		panic("Flashy octo is not charged up!")
	}

	theOne.RecentlyFlashed = true
	theOne.PowerLevel = 0
	d.flashCount++
	d.lastFlashCount++

	d.grid[point.X][point.Y] = theOne

	var nextFlashies = []Point{}
	for _, direction := range helpers.EightDirections {
		neiX := point.X + direction.X
		neiY := point.Y + direction.Y

		if neiX < 0 || neiY < 0 || neiX >= len(d.grid) || neiY >= len(d.grid[neiX]) {
			continue
		}

		neighbour := d.grid[neiX][neiY]

		if neighbour.RecentlyFlashed {
			continue
		}

		neighbour.PowerLevel += 1

		if neighbour.PowerLevel > 9 {
			nextFlashies = append(nextFlashies, Point{neiX, neiY})
		}

		d.grid[neiX][neiY] = neighbour
	}

	return nextFlashies
}

func (d DumboSimulator) Dump() {
	for x := 0; x < len(d.grid); x++ {
		for y := 0; y < len(d.grid[x]); y++ {
			fmt.Print(d.grid[x][y].PowerLevel)
		}
		fmt.Println()
	}
	fmt.Println("==========================")
}

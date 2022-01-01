package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"strings"
)

func Day5Vents() error {
	var vents []VentCoordinate
	helpers.ReadLines("inputs/day5.txt", func(s string) error {
		vents = append(vents, ParseVentCoordinates(s))

		return nil
	})

	plot := PlotVents(CardinalVentsOnly(vents))
	overlaps := CountVentOverlaps(plot)

	fmt.Println("Day 5 - Part 1:", overlaps)

	plot = PlotVents(vents)
	overlaps = CountVentOverlaps(plot)

	fmt.Println("Day 5 - Part 2:", overlaps)

	return nil
}

type VentCoordinate struct {
	X1, Y1 int
	X2, Y2 int
}

func (c VentCoordinate) Cardinal() bool {
	return c.X1 == c.X2 || c.Y1 == c.Y2
}

func (c VentCoordinate) Direction() helpers.Vec2 {
	if c.X1 == c.X2 {
		if c.Y1 < c.Y2 {
			return helpers.DownVector
		} else {
			return helpers.UpVector
		}
	}

	if c.Y1 == c.Y2 {
		if c.X1 < c.X2 {
			return helpers.RightVector
		} else {
			return helpers.LeftVector
		}
	}

	if c.Y1 < c.Y2 {
		if c.X1 < c.X2 {
			return helpers.DownRightVector
		} else {
			return helpers.DownLeftVector
		}
	} else {
		if c.X1 < c.X2 {
			return helpers.UpRightVector
		} else {
			return helpers.UpLeftVector
		}
	}
}

func ParseVentCoordinates(line string) VentCoordinate {
	ints := helpers.StringsToInts(strings.Split(strings.ReplaceAll(line, " -> ", ","), ","))

	return VentCoordinate{
		X1: ints[0],
		Y1: ints[1],
		X2: ints[2],
		Y2: ints[3],
	}
}

func CardinalVentsOnly(vents []VentCoordinate) []VentCoordinate {
	var cardinals []VentCoordinate
	for _, v := range vents {
		if v.Cardinal() {
			cardinals = append(cardinals, v)
		}
	}
	return cardinals
}

func PlotVents(vents []VentCoordinate) [][]int {
	sizeX, sizeY := VentPlotDimensionFromCoordinates(vents)

	plot := make([][]int, sizeX+1)
	for i := range plot {
		plot[i] = make([]int, sizeY+1)
	}

	for _, vent := range vents {
		start := helpers.Vec2{X: vent.X1, Y: vent.Y1}
		end := helpers.Vec2{X: vent.X2, Y: vent.Y2}
		direction := vent.Direction()

		for !start.Equal(end) {
			plot[start.X][start.Y] += 1
			start = start.Add(direction)
		}

		plot[start.X][start.Y] += 1
	}

	return plot
}

func PrintVentPlot(plot [][]int) {
	for y := 0; y < len(plot[0]); y++ {
		for x := 0; x < len(plot); x++ {
			cell := plot[x][y]

			if cell == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", cell)
			}
		}
		fmt.Println()
	}
}
func CountVentOverlaps(plot [][]int) int {
	var overlaps = 0
	for _, row := range plot {
		for _, v := range row {
			if v > 1 {
				overlaps += 1
			}
		}
	}
	return overlaps
}

func VentPlotDimensionFromCoordinates(vents []VentCoordinate) (int, int) {
	var x, y int
	for _, v := range vents {
		if v.X1 > x {
			x = v.X1
		}
		if v.X2 > x {
			x = v.X2
		}
		if v.Y1 > y {
			y = v.Y1
		}
		if v.Y2 > y {
			y = v.Y2
		}
	}

	return x, y
}

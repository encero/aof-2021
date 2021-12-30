package solutions

import (
	"github.com/encero/advent-of-code-2021/helpers"
	is_ "github.com/matryer/is"
	"testing"
)

func TestParseVentCoordinates(t *testing.T) {
	is := is_.New(t)

	coords := ParseVentCoordinates("242,601 -> 242,18")

	is.Equal(coords.X1, 242)
	is.Equal(coords.Y1, 601)
	is.Equal(coords.X2, 242)
	is.Equal(coords.Y2, 18)

	is.True(coords.Cardinal())

	coords = ParseVentCoordinates("8,0 -> 0,8")
	is.Equal(coords.X1, 8)
	is.Equal(coords.Y1, 0)
	is.Equal(coords.X2, 0)
	is.Equal(coords.Y2, 8)

	is.True(!coords.Cardinal())

	// 0,0 --------- 1,0
	// |             |
	// |             |
	// 1,0 --------- 0,1
	coords = ParseVentCoordinates("0,0 -> 1,0")
	is.Equal(coords.Direction(), helpers.RightVector) // up

	coords = ParseVentCoordinates("0,0 -> 0,1")
	is.Equal(coords.Direction(), helpers.DownVector) // down
}

var vents = []VentCoordinate{
	{0, 9, 5, 9},
	{8, 0, 0, 8},
	{9, 4, 3, 4},
	{2, 2, 2, 1},
	{7, 0, 7, 4},
	{6, 4, 2, 0},
	{0, 9, 2, 9},
	{3, 4, 1, 4},
	{0, 0, 8, 8},
	{5, 5, 8, 2},
}

func TestPlotVents(t *testing.T) {
	is := is_.New(t)

	cardinals := CardinalVentsOnly(vents)

	plot := PlotVents(cardinals)
	//PrintVentPlot(plot)

	overlap := CountVentOverlaps(plot)

	is.Equal(overlap, 5) // test overlap without diagonals

	plot = PlotVents(vents)
	//	PrintVentPlot(plot)

	overlap = CountVentOverlaps(plot)

	is.Equal(overlap, 12) // test overlap
}

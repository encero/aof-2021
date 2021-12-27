package solutions

import (
	"testing"

	is_ "github.com/matryer/is"
)

var chitonsGrid = [][]int{
	{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
	{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
	{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
	{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
	{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
	{1, 3, 1, 9, 1, 2, 8, 1, 3, 7},
	{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
	{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
	{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
	{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
}

func TestChitonPath(t *testing.T) {
	is := is_.New(t)

	risk := ChitonPath(chitonsGrid)

	is.Equal(risk, 40)

	//helpers.PrintIntGrid(ExpandChitonGrid(chitonsGrid))

	risk = ChitonPath(ExpandChitonGrid(chitonsGrid))

	is.Equal(risk, 315)
}

func TestExpandChitonGrid(t *testing.T) {
	is := is_.New(t)

	grid := [][]int{{8}}

	grid = ExpandChitonGrid(grid)

	is.Equal(grid, [][]int{
		{8, 9, 1, 2, 3},
		{9, 1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
	})
}

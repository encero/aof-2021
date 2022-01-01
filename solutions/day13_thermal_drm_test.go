package solutions

import (
	"github.com/encero/advent-of-code-2021/helpers"
	is_ "github.com/matryer/is"
	"testing"
)

var testPaperDots = []helpers.Vec2{
	{6, 10},
	{0, 14},
	{9, 10},
	{0, 3},
	{10, 4},
	{4, 11},
	{6, 0},
	{6, 12},
	{4, 1},
	{0, 13},
	{10, 12},
	{3, 4},
	{3, 0},
	{8, 4},
	{1, 10},
	{2, 14},
	{8, 10},
	{9, 0},
}

func TestPaperFold(t *testing.T) {
	is := is_.New(t)

	paper := NewPaper(testPaperDots)

	paper = paper.Fold(FoldY, 7)

	is.Equal(paper.Dots(), 17) // first fold

	paper = paper.Fold(FoldX, 5)

	is.Equal(paper.Dots(), 16) // second fold
}

package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"strings"
)


func Day13ThermalDrm() error {
	var dots []helpers.Vector

	err := helpers.ReadLines("inputs/day13_dots.txt", func(line string) error {
		coords := helpers.StringsToInts(strings.Split(line, ","))

		dots = append(dots, helpers.Vector{X: coords[0], Y: coords[1]})

		return nil
	})
	if err != nil {
		return err
	}

	type fold struct {
		dir FoldDirection
		pos int
	}

	folds := []fold{
		{FoldX, 655},
		{FoldY, 447},
		{FoldX, 327},
		{FoldY, 223},
		{FoldX, 163},
		{FoldY, 111},
		{FoldX, 81},
		{FoldY, 55},
		{FoldX, 40},
		{FoldY, 27},
		{FoldY, 13},
		{FoldY, 6},
	}

	paper := NewPaper(dots)

	paper = paper.Fold(folds[0].dir, folds[0].pos)

	fmt.Println("Part 1:", paper.Dots())

	for i := 1; i < len(folds); i++ {
		paper = paper.Fold(folds[i].dir, folds[i].pos)
	}

	paper.Dump()

	return nil
}

type FoldDirection int

const (
	FoldX FoldDirection = iota
	FoldY
)

type Paper struct {
	dots map[helpers.Vector]struct{}
}

func NewPaper(dots []helpers.Vector) Paper {
	dotMap := make(map[helpers.Vector]struct{}, len(dots))

	for _, dot := range dots {
		dotMap[dot] = struct{}{}
	}

	return Paper{
		dots: dotMap,
	}
}

func (p Paper) Fold(dir FoldDirection, pos int) Paper {
	foldXFunc := func(dot helpers.Vector) (*helpers.Vector, bool) {
		if dot.X > pos {
			return &helpers.Vector{X: pos - (dot.X - pos), Y: dot.Y}, false
		}

		if dot.X == pos {
			return nil, false
		}

		return &helpers.Vector{X: dot.X, Y: dot.Y}, true
	}

	foldYFunc := func(dot helpers.Vector) (*helpers.Vector, bool) {
		if dot.Y > pos {
			return &helpers.Vector{X: dot.X, Y: pos - (dot.Y - pos)}, false
		}

		if dot.Y == pos {
			return nil, false
		}

		return &helpers.Vector{X: dot.X, Y: dot.Y}, true
	}

	foldF := foldXFunc
	if dir == FoldY {
		foldF = foldYFunc
	}

	var dotMap = make(map[helpers.Vector]struct{}, len(p.dots))

	for k := range p.dots {
		folded, _ := foldF(k)

		if folded != nil {
			dotMap[*folded] = struct{}{}
		}
	}

	return Paper{
		dots: dotMap,
	}
}

func (p Paper) Dots() int {
	return len(p.dots)
}

func (p Paper) Dump() {
	var maxX, maxY int
	for dot := range p.dots {
		fmt.Printf("%d,%d\n", dot.X, dot.Y)
		if dot.X > maxX {
			maxX = dot.X
		}

		if dot.Y > maxY {
			maxY = dot.Y
		}
	}

	for y := 0; y <= maxY; y++ {
		fmt.Printf("%2d ", y)
		for x := 0; x <= maxX; x++ {
			if _, ok := p.dots[helpers.Vector{X: x, Y: y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

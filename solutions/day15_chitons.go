package solutions

import (
	"fmt"

	"github.com/encero/advent-of-code-2021/helpers"
)

func Day15Chitons() error {
	grid := helpers.ReadIntGrid("inputs/day15.txt")

	finder := NewChitonsPathFinder(grid)
	risk := finder.FindPath()

	fmt.Println("day 15, Part 1:", risk)

    finder = NewChitonsPathFinder(ExpandChitonGrid(grid))
    risk = finder.FindPath()

    fmt.Println("day 15, Part 2:", risk)
	return nil
}

func ChitonPath(grid [][]int) int {
	finder := NewChitonsPathFinder(grid)
	risk := finder.FindPath()

	return risk
}

type ChitonsPathFinder struct {
	chitons [][]int
	open    map[helpers.Vector]struct{}
	gScore  map[helpers.Vector]int
	end     helpers.Vector
}

func NewChitonsPathFinder(chitons [][]int) *ChitonsPathFinder {
	start := helpers.Vector{X: 0, Y: 0}
	open := make(map[helpers.Vector]struct{})
	open[start] = struct{}{}

	end := helpers.Vector{X: len(chitons) - 1, Y: len(chitons[0]) - 1}

	gScore := make(map[helpers.Vector]int)
	gScore[start] = 0

	return &ChitonsPathFinder{
		chitons: chitons,
		open:    open,
		gScore:  gScore,
		end:     end,
	}
}

func (ch *ChitonsPathFinder) FindPath() int {
	for len(ch.open) > 0 {
		current := ch.cheapestOpen()

		//fmt.Printf("%02d:%02d %d->%d\n", current.X, current.Y, ch.gScore[current], current.ManhattanDistance(ch.end))

		if current.Equal(ch.end) {
			return ch.gScore[current]
		}

		delete(ch.open, current)
		for _, dir := range helpers.FourDirections {
			nei := current.Add(dir)

			if nei.X < 0 || nei.Y < 0 || nei.X > ch.end.X || nei.Y > ch.end.Y {
				continue
			}

			score := ch.gScore[current] + ch.chitons[nei.X][nei.Y]
			if neiScore, ok := ch.gScore[nei]; !ok || score < neiScore {
				ch.gScore[nei] = score
				if _, ok := ch.open[nei]; !ok {
					ch.open[nei] = struct{}{}
				}
			}
		}
	}

	panic("no solution")
}

func (ch *ChitonsPathFinder) cheapestOpen() helpers.Vector {
	cost := 12345678
	var cheapest helpers.Vector

	for v := range ch.open {
		if newCost := v.ManhattanDistance(ch.end) + ch.gScore[v]; newCost < cost {
			cost = newCost
			cheapest = v
		}
	}

	return cheapest
}

func ExpandChitonGrid(input [][]int) [][]int {
    expanded := make([][]int, len(input)*5)

    for x := 0; x < len(input) * 5; x++ {
        expanded[x] = make([]int, len(input[0]) * 5)
        for y := 0; y < len(input[0]) * 5; y ++ {
            inputX := x % len(input)
            inputY := y % len(input[0])


            coef := helpers.Vector{X:x/len(input), Y: y/len(input[0])}.ManhattanDistance(helpers.ZeroVector)

            //fmt.Println(x,y, "->", inputX, inputY, "c", coef)
            expanded[x][y] = input[inputX][inputY] + coef
            if expanded[x][y] > 9 {
                expanded[x][y] -= 9
            }
        }
    }

    return expanded
}

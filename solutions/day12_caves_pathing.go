package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"strings"
)

func Day12CavesPathing() error {
	lines, err := helpers.ReadAllLines("inputs/day12.txt")
	if err != nil {
		return err
	}

	caves := MapCaves(lines)

	paths := NewCaveNavigator(caves, Part1CaveFunction).Solve()

	fmt.Println("Day 12: Part 1:", len(paths))

	paths = NewCaveNavigator(caves, Part2CaveFunction).Solve()

	fmt.Println("Day 12: Part 2:", len(paths))

	return nil
}

type Cave struct {
	ID     string
	linked []string
}

type CavePath struct {
	Path    []string
	Doubled bool
}

func (c Cave) Linked() []string {
	return c.linked
}

func (c Cave) IsBig() bool {
	return strings.ToUpper(c.ID) == c.ID
}

func (c Cave) IsSmall() bool {
	return !c.IsBig()
}

type CaveNavigator struct {
	caves    map[string]*Cave
	caveFunc func(*Cave, *CavePath) bool
}

func Part1CaveFunction(cave *Cave, path *CavePath) bool {
	return cave.IsSmall() && helpers.StringSliceContains(path.Path, cave.ID)
}

func Part2CaveFunction(cave *Cave, path *CavePath) bool {
	var count int

	if cave.ID == "start" {
		return true
	}

	if cave.IsBig() {
		return false
	}

	for _, node := range path.Path {
		if node == cave.ID {
			count++
		}
	}

	if count == 0 {
		return false
	}

	if count == 1 && !path.Doubled {
		path.Doubled = true
		return false
	}

	return true
}

func NewCaveNavigator(caves map[string]*Cave, caveFunc func(*Cave, *CavePath) bool) *CaveNavigator {
	if caveFunc == nil {
		caveFunc = Part1CaveFunction
	}

	return &CaveNavigator{
		caves:    caves,
		caveFunc: caveFunc,
	}
}

func (n CaveNavigator) Solve() []*CavePath {

	return n.solve(&CavePath{
		Path: []string{"start"},
	})
}

func (n CaveNavigator) solve(path *CavePath) []*CavePath {
	lastCave := path.Path[len(path.Path)-1]

	if lastCave == "end" {
		return []*CavePath{path}
	}

	var out []*CavePath

	branches := n.caves[lastCave].Linked()
	for _, branchID := range branches {
		branch := n.caves[branchID]

		newPath := &CavePath{
			Path:    append(make([]string, 0, len(path.Path)+1), path.Path...),
			Doubled: path.Doubled,
		}

		if n.caveFunc(branch, newPath) {
			continue
		}

		newPath.Path = append(newPath.Path, branchID)

		newPaths := n.solve(newPath)

		out = append(out, newPaths...)
	}

	return out
}

func MapCaves(input []string) map[string]*Cave {
	out := make(map[string]*Cave)

	link := func(A, B string) {
		cave, ok := out[A]
		if !ok {
			cave = &Cave{ID: A, linked: []string{B}}
			out[A] = cave
		} else {
			cave.linked = append(cave.linked, B)
		}
	}

	for _, line := range input {
		parts := strings.Split(line, "-")
		A, B := parts[0], parts[1]

		link(A, B)
		link(B, A)
	}

	return out
}

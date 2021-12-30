package solutions

import (
	"fmt"
	"strings"

	"github.com/encero/advent-of-code-2021/helpers"
)

func Day19Beacons() error {
	lines, err := helpers.ReadAllLines("inputs/day19.txt")
	if err != nil {
		return err
	}

	scanners := ParseScannerInput(lines)
	MatchScanners(scanners)

	uni := UniqueBeacons(scanners)

	fmt.Println("Unique beacons:", len(uni))
	fmt.Println("max distance:", BigestManhattan(scanners))
	return nil
}

type Scanner struct {
	Beacons         []helpers.Vec3
	Rotation        helpers.Vec3
	Offset          helpers.Vec3
	GlobalOffset    helpers.Vec3
	MatchingScanner int
}

func (s Scanner) Get(i int) helpers.Vec3 {
	return s.Beacons[i].
		RotateX(s.Rotation.X).
		RotateY(s.Rotation.Y).
		RotateZ(s.Rotation.Z).
		Add(s.Offset)
}

func (s Scanner) GetGlobal(i int) helpers.Vec3 {
	return s.Beacons[i].
		RotateX(s.Rotation.X).
		RotateY(s.Rotation.Y).
		RotateZ(s.Rotation.Z).
		Add(s.GlobalOffset)
}

func (s *Scanner) MatchOffset(b *Scanner, ai, bi int) {
	s.Offset = helpers.Vec3{}

	aV := s.Get(ai)
	bV := b.Get(bi)

	s.Offset = helpers.Vec3{
		X: bV.X - aV.X,
		Y: bV.Y - aV.Y,
		Z: bV.Z - aV.Z,
	}
}

func ParseScannerInput(in []string) []*Scanner {
	scanners := []*Scanner{}

	for _, line := range in {
		if line == "" {
			continue
		}
		if string(line[0:3]) == "---" {
			scanners = append(scanners, &Scanner{})
			continue
		}

		coords := helpers.StringsToInts(strings.Split(line, ","))

		scanners[len(scanners)-1].Beacons = append(scanners[len(scanners)-1].Beacons, helpers.Vec3{X: coords[0], Y: coords[1], Z: coords[2]})
	}

	return scanners
}

func MatchedLocations(a, b *Scanner) map[helpers.Vec3]struct{} {
	matched := make(map[helpers.Vec3]struct{})
	aMap := make(map[helpers.Vec3]struct{})

	for i := range a.Beacons {
		aMap[a.Get(i)] = struct{}{}
	}

	for i := range b.Beacons {
		v := b.Get(i)
		if _, ok := aMap[v]; ok {
			matched[v] = struct{}{}
		}
	}

	return matched
}

func MatchTwoScanners(a, b *Scanner) map[helpers.Vec3]struct{} {
	for _, rot := range helpers.Rotations {
		b.Rotation = rot
		for ai := 0; ai < len(a.Beacons); ai++ {
			for bi := ai + 1; bi < len(b.Beacons); bi++ {
				b.MatchOffset(a, bi, ai)

				matched := MatchedLocations(a, b)
				if len(matched) >= 12 {

					b.GlobalOffset = a.GlobalOffset.Add(b.Offset)
					return matched
				}
			}
		}
	}

	return nil
}

func MatchScanners(scanners []*Scanner) {
	matched := map[int]struct{}{
		0: {},
	}

	count := 1000
	for len(matched) < len(scanners) && count > 0 {
		count--

		for i, s := range scanners {
			if _, ok := matched[i]; ok {
				continue
			}

			for m := range matched {
				ok := MatchTwoScanners(scanners[m], s)
				if ok != nil {
					fmt.Println("matched", m, "to", i, "with", len(ok), "matches")
					matched[i] = struct{}{}
					s.MatchingScanner = m
					break
				}
			}
		}
	}

	if len(matched) < len(scanners) {
		panic("match not found")
	}
}

func UniqueBeacons(scanners []*Scanner) map[helpers.Vec3]struct{} {
	uni := make(map[helpers.Vec3]struct{})

	for si, s := range scanners {
		fmt.Println(si, "off", s.Offset, "global", s.GlobalOffset, "match", s.MatchingScanner)
		for i := range s.Beacons {
			//uni[s.GetGlobal(i)] = struct{}{}
			uni[s.Get(i)] = struct{}{}
		}
	}

	return uni
}

func BigestManhattan(scanners []*Scanner) int {
	max := 0

	for _, s := range scanners {
		for _, s2 := range scanners {
			distance := Vec3Manhattan(s.Offset, s2.Offset)
			if distance > max {
				max = distance
			}
		}
	}

	return max
}

func Vec3Manhattan(a, b helpers.Vec3) int {
	distance := 0
	dif := func(a, b int) int {
		if a > b {
			return a - b
		} else {
			return b - a
		}
	}

	distance += dif(a.X, b.X)
	distance += dif(a.Y, b.Y)
	distance += dif(a.Z, b.Z)

	return distance
}

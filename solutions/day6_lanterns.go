package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"strings"
)

func Day6Lanterns() error {
	initial := helpers.StringsToInts(strings.Split(helpers.ReadLine("inputs/day6.txt"), ","))

	colony := NewLanternColony(initial)

	for i := 0; i < 80; i++ {
		colony.Evolve()
	}

	fmt.Println("Day 6 Part 1", colony.Population())

	colony = NewLanternColony(initial)

	for i := 0; i < 256; i++ {
		colony.Evolve()
	}

	fmt.Println("Day 6 Part 2", colony.Population())

	return nil
}

type LanternColony struct {
	day     int
	types   []int
	waiting []int
}

func NewLanternColony(initial []int) *LanternColony {
	types := make([]int, 7)

	for _, t := range initial {
		types[t]++
	}

	return &LanternColony{
		types:   types,
		waiting: make([]int, 2),
	}
}

func (c *LanternColony) Evolve() {
	effectiveDay := c.day % 7

	spawn := c.types[effectiveDay]

	c.waiting = append(c.waiting, spawn)
	c.types[effectiveDay] += c.waiting[0]
	c.waiting = c.waiting[1:]

	c.day += 1
}

func (c LanternColony) Population() int {
	return helpers.SliceSum(c.types) + helpers.SliceSum(c.waiting)
}

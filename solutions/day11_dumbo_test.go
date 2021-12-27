package solutions

import (
	is_ "github.com/matryer/is"
	"testing"
)

var testDumboGrid = [][]DumboOctopus{
	{{PowerLevel: 5}, {PowerLevel: 4}, {PowerLevel: 8}, {PowerLevel: 3}, {PowerLevel: 1}, {PowerLevel: 4}, {PowerLevel: 3}, {PowerLevel: 2}, {PowerLevel: 2}, {PowerLevel: 3}},
	{{PowerLevel: 2}, {PowerLevel: 7}, {PowerLevel: 4}, {PowerLevel: 5}, {PowerLevel: 8}, {PowerLevel: 5}, {PowerLevel: 4}, {PowerLevel: 7}, {PowerLevel: 1}, {PowerLevel: 1}},
	{{PowerLevel: 5}, {PowerLevel: 2}, {PowerLevel: 6}, {PowerLevel: 4}, {PowerLevel: 5}, {PowerLevel: 5}, {PowerLevel: 6}, {PowerLevel: 1}, {PowerLevel: 7}, {PowerLevel: 3}},
	{{PowerLevel: 6}, {PowerLevel: 1}, {PowerLevel: 4}, {PowerLevel: 1}, {PowerLevel: 3}, {PowerLevel: 3}, {PowerLevel: 6}, {PowerLevel: 1}, {PowerLevel: 4}, {PowerLevel: 6}},
	{{PowerLevel: 6}, {PowerLevel: 3}, {PowerLevel: 5}, {PowerLevel: 7}, {PowerLevel: 3}, {PowerLevel: 8}, {PowerLevel: 5}, {PowerLevel: 4}, {PowerLevel: 7}, {PowerLevel: 8}},
	{{PowerLevel: 4}, {PowerLevel: 1}, {PowerLevel: 6}, {PowerLevel: 7}, {PowerLevel: 5}, {PowerLevel: 2}, {PowerLevel: 4}, {PowerLevel: 6}, {PowerLevel: 4}, {PowerLevel: 5}},
	{{PowerLevel: 2}, {PowerLevel: 1}, {PowerLevel: 7}, {PowerLevel: 6}, {PowerLevel: 8}, {PowerLevel: 4}, {PowerLevel: 1}, {PowerLevel: 7}, {PowerLevel: 2}, {PowerLevel: 1}},
	{{PowerLevel: 6}, {PowerLevel: 8}, {PowerLevel: 8}, {PowerLevel: 2}, {PowerLevel: 8}, {PowerLevel: 8}, {PowerLevel: 1}, {PowerLevel: 1}, {PowerLevel: 3}, {PowerLevel: 4}},
	{{PowerLevel: 4}, {PowerLevel: 8}, {PowerLevel: 4}, {PowerLevel: 6}, {PowerLevel: 8}, {PowerLevel: 4}, {PowerLevel: 8}, {PowerLevel: 5}, {PowerLevel: 5}, {PowerLevel: 4}},
	{{PowerLevel: 5}, {PowerLevel: 2}, {PowerLevel: 8}, {PowerLevel: 3}, {PowerLevel: 7}, {PowerLevel: 5}, {PowerLevel: 1}, {PowerLevel: 5}, {PowerLevel: 2}, {PowerLevel: 6}},
}

func TestDumboSimulator(t *testing.T) {
	is := is_.New(t)

	simulator := NewDumboSimulator(testDumboGrid)

	for i := 0; i < 10; i++ {
		//simulator.Dump()
		simulator.Step()
	}

	is.Equal(simulator.Flashes(), 204) // flash count

	for i := 10; i < 100; i++ {
		simulator.Step()
	}

	is.Equal(simulator.Flashes(), 1656) // flash count
}

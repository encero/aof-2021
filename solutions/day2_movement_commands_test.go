package solutions

import "testing"
import is_ "github.com/matryer/is"


func TestSimulateMovement(t *testing.T ) {
	is := is_.New(t)

	depth, x := SimulateMovement([]MovementCommand{
		{Forward, 1},
		{Forward, 2},
		{Forward, 3},
		{Forward, 4},
	})
	is.Equal(depth, 0)
	is.Equal(x, 10)

	depth, x = SimulateMovement([]MovementCommand{
		{Down, 2},
		{Forward, 2},
		{Up, 1},
		{Down, 2},
	})

	is.Equal(depth, 3)
	is.Equal(x, 2)
}

func TestSimulateAimMovement(t *testing.T) {
	is := is_.New(t)

	depth, x := SimulateAimMovement([]MovementCommand{
		{Forward, 1},
		{Forward, 2},
		{Forward, 3},
		{Forward, 4},
	})
	is.Equal(depth, 0)
	is.Equal(x, 10)

	depth, x = SimulateAimMovement([]MovementCommand{
		{Forward, 5},
		{Down, 5},
		{Forward, 8},
		{Up, 3},
		{Down, 8},
		{Forward, 2},
	})

	is.Equal(depth, 60)
	is.Equal(x, 15)
}

func TestParseDirection(t *testing.T) {
	is := is_.New(t)

	is.Equal(parseDirection("up"), Up) // Up
	is.Equal(parseDirection("down"), Down) // Down
	is.Equal(parseDirection("forward"), Forward) // Forward
}

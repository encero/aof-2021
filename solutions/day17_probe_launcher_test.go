package solutions

import (
	"testing"

	"github.com/encero/advent-of-code-2021/helpers"
	is_ "github.com/matryer/is"
)

func TestProbeMove(t *testing.T) {
	is := is_.New(t)

	probe := Probe{
		Velocity: helpers.Vec2{X: 1, Y: 1},
	}

	probe.Move()
	is.Equal(probe.Position, helpers.Vec2{X: 1, Y: 1})
	is.Equal(probe.Velocity, helpers.Vec2{X: 0, Y: 0})

	probe.Move()
	is.Equal(probe.Position, helpers.Vec2{X: 1, Y: 1})
	is.Equal(probe.Velocity, helpers.Vec2{X: 0, Y: -1})
}

func TestBounds(t *testing.T) {
	is := is_.New(t)

	b := Bounds{
		10, 20,
		-30, -40,
	}

	is.Equal(b.Hit(helpers.Vec2{0, 0}), false)   // Hit out
	is.Equal(b.Hit(helpers.Vec2{10, -30}), true) // Hit in

	is.Equal(b.Over(helpers.Vec2{30, -50}), true) // Over
	is.Equal(b.Over(helpers.Vec2{0, 0}), false)   // Not Over
}

func TestHitTarget(t *testing.T) {
	is := is_.New(t)
	bounds := Bounds{20, 30, -5, -10}

	probe := &Probe{
		Velocity: helpers.Vec2{X: 6, Y: 9},
	}

	wasHit := probe.HitTarget(bounds)

	is.Equal(probe.MaxY, 45) // maxY
	is.Equal(wasHit, true)   // was hit
}

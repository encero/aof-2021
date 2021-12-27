package solutions

import (
	"fmt"

	"github.com/encero/advent-of-code-2021/helpers"
)

func Day17ProbeLauncher() error {
    maxY, hits := OptimizeProbeHeight()

    fmt.Println("Day 17, Part 1:", maxY)
    fmt.Println("Day 17, Part 2:", hits)


    return nil
}

func OptimizeProbeHeight() (int, int) {
    // 236..262, y=-78..-58
    bounds := Bounds{236, 262, -58, -78}

    maxY := 0
    hits := 0

    for x:= 0; x <= 262; x++ {
        for y := -90; y <= 300; y++ {
            probe := Probe{
                Velocity: helpers.Vector{x, y},
            }

            hit := probe.HitTarget(bounds)
            if hit && probe.MaxY > maxY {
                maxY = probe.MaxY
            }
            if hit {
                hits ++
            }
        }
    }

    return maxY, hits
}

type Probe struct {
	Velocity helpers.Vector
	Position helpers.Vector
	MaxY     int
}

func (p *Probe) Move() {
	p.Position = p.Position.Add(p.Velocity)
	p.Velocity.Y -= 1

	if p.Velocity.X < 0 {
		p.Velocity.X += 1
	} else if p.Velocity.X > 0 {
		p.Velocity.X -= 1
	}

	if p.MaxY < p.Position.Y {
		p.MaxY = p.Position.Y
	}

    //fmt.Println(p.Position.X, p.Position.Y)
}

func (p *Probe) HitTarget(bounds Bounds) bool {
	for !bounds.Hit(p.Position) && !bounds.Over(p.Position) {
		p.Move()
	}

    //fmt.Println("hit", bounds.Hit(p.Position))
    //fmt.Println("over", bounds.Over(p.Position))

	return bounds.Hit(p.Position)
}

type Bounds struct {
	X0, X1 int
	Y0, Y1 int
}

func (b *Bounds) Hit(v helpers.Vector) bool {
	return v.X >= b.X0 && v.X <= b.X1 && v.Y <= b.Y0 && v.Y >= b.Y1
}

func (b *Bounds) Over(v helpers.Vector) bool {
	return v.X >= b.X1 || v.Y <= b.Y1
}

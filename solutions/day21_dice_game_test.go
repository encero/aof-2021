package solutions

import (
	"testing"

	is_ "github.com/matryer/is"
)

func TestDiracGamePlay(t *testing.T){
    if testing.Short() {
        t.Skip()
    }

    is := is_.New(t)

    g := Game{
        p1: &Player{
            Position:4,
        },
        p2: &Player{
            Position: 8,
        },
        UniverseDepth: 1,
    }

    PlayRecursive(&g)

    is.Equal(g.p1.Wins, int64(444356092776315))
    is.Equal(g.p2.Wins, int64(341960390180808))

}
func TestGamePlay(t *testing.T) {
    is := is_.New(t)

    g := Game{
        p1: &Player{
            Position:4,
        },
        p2: &Player{
            Position: 8,
        },
        Die: &DeterministicDice{},
    }

    winner := g.Play()

    is.Equal(winner, g.p1)
    is.Equal(g.p2.Points, 745)
    is.Equal(g.Die.RollCount(), 993)
}
func TestDeterministicDice(t *testing.T) {
    is := is_.New(t)

    dice := DeterministicDice{}

    for i := 1; i <= 100; i++ {
        is.Equal(dice.Roll(), i) // dice has deterministic output
    }

    is.Equal(dice.Roll(), 1) // dice will roll over

    is.Equal(dice.RollCount(), 101)
}

func TestPlayerAdvance(t *testing.T) {
    is := is_.New(t)

    p := Player{
        Position: 4,
    }

    p.Advance(2)
    is.Equal(p.Points, 6)
    is.Equal(p.Position, 6)

    p.Advance(4)
    is.Equal(p.Points, 16)
    is.Equal(p.Position, 10)

    
    p.Advance(1)
    is.Equal(p.Points, 17)
    is.Equal(p.Position, 1)

    p.Advance(50)
    is.Equal(p.Points, 18)
    is.Equal(p.Position, 1)
}

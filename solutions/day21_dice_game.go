package solutions

import "fmt"

func Day21DiceGame() error {
	game := Game{
		p1: &Player{
			Position: 6,
		},
		p2: &Player{
			Position: 3,
		},
		Die: &DeterministicDice{},
	}

	winner := game.Play()

	loser := game.p1
	if winner == game.p1 {
		loser = game.p2
	}

	fmt.Println("Dice game result:", loser.Points*game.Die.RollCount())

	game = Game{
		p1: &Player{
			Position: 6,
		},
		p2: &Player{
			Position: 3,
		},
        UniverseDepth: 1,
	}

    PlayRecursive(&game)

    if game.p1.Wins > game.p2.Wins {
        fmt.Println("most universes wins have p1 with:", game.p1.Wins)
    } else {
        fmt.Println("most universes wins have p2 with:", game.p2.Wins)
    }

	return nil
}

type Game struct {
	p1            *Player
	p2            *Player
	Die           Dice
	UniverseDepth int64
}

func (g *Game) Play() *Player {
	roll := func() int {
		return g.Die.Roll() + g.Die.Roll() + g.Die.Roll()
	}

	for {
		amount := roll()

		g.p1.Advance(amount)
		//fmt.Printf("P1 rolls: %d moves to: %d total score: %d\n", amount, g.p1.Position, g.p1.Points)
		if g.p1.Points >= 1000 {
			return g.p1
		}

		amount = roll()
		g.p2.Advance(amount)
		//fmt.Printf("P2 rolls: %d moves to: %d total score: %d\n", amount, g.p2.Position, g.p2.Points)
		if g.p2.Points >= 1000 {
			return g.p2
		}
	}
}

var diracDiceOutcome = []int{3, 4, 5, 6, 7, 8, 9}
var diracDiceUniverses = []int{1, 3, 6, 7, 6, 3, 1}

func PlayRecursive(g *Game) {
	for p1o := range diracDiceOutcome {
		ng := Game{
			p1: &Player{
				Position: g.p1.Position,
				Points:   g.p1.Points,
			},
			p2: &Player{
				Position: g.p2.Position,
				Points:   g.p2.Points,
			},
			UniverseDepth: g.UniverseDepth,
		}

		ng.p1.Advance(diracDiceOutcome[p1o])
		ng.UniverseDepth *= int64(diracDiceUniverses[p1o])
		if ng.p1.Points >= 21 {
			g.p1.Wins += ng.UniverseDepth
			continue
		}

		for p2o := range diracDiceOutcome {
			ng2 := Game{
				p1: &Player{
					Position: ng.p1.Position,
					Points:   ng.p1.Points,
				},
				p2: &Player{
					Position: g.p2.Position,
					Points:   g.p2.Points,
				},
                UniverseDepth: ng.UniverseDepth,
			}

			ng2.p2.Advance(diracDiceOutcome[p2o])
			ng2.UniverseDepth *= int64(diracDiceUniverses[p2o])
			if ng2.p2.Points >= 21 {
				g.p2.Wins += ng2.UniverseDepth
				continue
			}

			PlayRecursive(&ng2)

			g.p1.Wins += ng2.p1.Wins
			g.p2.Wins += ng2.p2.Wins
		}
	}
}

type Dice interface {
	Roll() int
	RollCount() int
}

type DeterministicDice struct {
	counter   int
	rollCount int
}

func (d *DeterministicDice) Roll() int {
	if d.counter == 100 {
		d.counter = 0
	}

	d.counter++
	d.rollCount++

	return d.counter
}

func (d *DeterministicDice) RollCount() int {
	return d.rollCount
}

type Player struct {
	Points   int
	Position int
	Wins     int64
}

func (p *Player) Advance(amount int) {
	p.Position += amount
	if p.Position > 10 {
		p.Position -= 10 * (p.Position / 10)
	}

	if p.Position == 0 {
		p.Position = 10
	}

	p.Points += p.Position
}

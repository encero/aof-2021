package solutions

import (
	is_ "github.com/matryer/is"
	"testing"
)

func TestEvolveLantern(t *testing.T) {
	is := is_.New(t)

	colony := NewLanternColony([]int{3, 4, 3, 1, 2})

	is.Equal(colony.Population(), 5) // initial population

	for i := 0; i < 18; i++ {
		colony.Evolve()
	}

	is.Equal(colony.Population(), 26) // after 18 generations

	for i := 18; i < 80; i++ {
		colony.Evolve()
	}

	is.Equal(colony.Population(), 5934) // after 80 generations
}

package solutions

import (
	"testing"

	is_ "github.com/matryer/is"
)

var polymerPairs = map[string]string{
	"CH": "B",
	"HH": "N",
	"CB": "H",
	"NH": "C",
	"HB": "C",
	"HC": "B",
	"HN": "C",
	"NN": "C",
	"BH": "H",
	"NC": "B",
	"NB": "B",
	"BN": "B",
	"BB": "N",
	"BC": "B",
	"CC": "N",
	"CN": "C",
}

func TestPolymerize(t *testing.T) {
	is := is_.New(t)
	polymer := "NNCB"

	polymer = Polymerize(polymer, polymerPairs)
	is.Equal(polymer, "NCNBCHB")

	polymer = Polymerize(polymer, polymerPairs)
	is.Equal(polymer, "NBCCNBBBCBHCB")

	polymer = Polymerize(polymer, polymerPairs)
	is.Equal(polymer, "NBBBCNCCNBBNBNBBCHBHHBCHB")

	polymer = Polymerize(polymer, polymerPairs)
	is.Equal(polymer, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB")
}

func TestPolymerize2(t *testing.T) {
	is := is_.New(t)

	polymer := PolymerToPairs("NNCB")

	polymer = Polymerize2(polymer, polymerPairs)
	is.Equal(polymer, PolymerToPairs("NCNBCHB"))

	polymer = Polymerize2(polymer, polymerPairs)
	is.Equal(polymer, PolymerToPairs("NBCCNBBBCBHCB"))

	polymer = Polymerize2(polymer, polymerPairs)
	is.Equal(polymer, PolymerToPairs("NBBBCNCCNBBNBNBBCHBHHBCHB"))

	polymer = Polymerize2(polymer, polymerPairs)
	is.Equal(polymer, PolymerToPairs("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"))

	polymer = PolymerToPairs("NNCB")

	for i := 0; i < 40; i++ {
		polymer = Polymerize2(polymer, polymerPairs)
	}
	diff := PolymerMinMax(polymer)
	is.Equal(diff, int64(2188189693529))

}

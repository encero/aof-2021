package solutions

import (
	"fmt"
	is_ "github.com/matryer/is"
	"sort"
	"testing"
)

var caveMappingInput = []string{
	"start-A",
	"start-b",
	"A-c",
	"A-b",
	"b-d",
	"A-end",
	"b-end",
}

func TestMapCaves(t *testing.T) {
	is := is_.New(t)

	caveMap := MapCaves(caveMappingInput)

	linked := caveMap["start"].Linked()
	sort.Strings(linked)
	is.Equal(linked, []string{"A", "b"})

	linked = caveMap["A"].Linked()
	sort.Strings(linked)
	is.Equal(linked, []string{"b", "c", "end", "start"})
}

func TestCave_IsBig(t *testing.T) {
	is := is_.New(t)

	cave := &Cave{
		ID: "A",
	}

	is.True(cave.IsBig())

	cave = &Cave{
		ID: "a",
	}

	is.True(!cave.IsBig())
}

func TestCaveNavigator(t *testing.T) {
	test := []struct {
		caves          map[string]*Cave
		expectedCount  int
		expectedCount2 int
	}{
		{
			caves:          MapCaves(caveMappingInput),
			expectedCount:  10,
			expectedCount2: 36,
		},
		{
			caves: MapCaves(
				[]string{"dc-end",
					"HN-start",
					"start-kj",
					"dc-start",
					"dc-HN",
					"LN-dc",
					"HN-end",
					"kj-sa",
					"kj-HN",
					"kj-dc"},
			),
			expectedCount:  19,
			expectedCount2: 103,
		},
		{
			caves: MapCaves(
				[]string{
					"fs-end",
					"he-DX",
					"fs-he",
					"start-DX",
					"pj-DX",
					"end-zg",
					"zg-sl",
					"zg-pj",
					"pj-he",
					"RW-he",
					"fs-DX",
					"pj-RW",
					"zg-RW",
					"start-pj",
					"he-WI",
					"zg-he",
					"pj-fs",
					"start-RW",
				},
			),
			expectedCount:  226,
			expectedCount2: 3509,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("#%d", tt.expectedCount), func(t *testing.T) {
			is := is_.New(t)

			var paths []*CavePath

			//paths := NewCaveNavigator(tt.caves, Part1CaveFunction).Solve()
			//is.Equal(len(paths), tt.expectedCount)

			paths = NewCaveNavigator(tt.caves, Part2CaveFunction).Solve()
			is.Equal(len(paths), tt.expectedCount2)
		})
	}
}

package solutions

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	is_ "github.com/matryer/is"
)

var ignoreParent = cmpopts.IgnoreFields(SnailNumber{}, "Parent")

func TestSplitSnailNumber(t *testing.T) {
	is := is_.New(t)

	is.Equal(SplitSnailNumber("[1,2]"), []string{"1", "2"})
	is.Equal(SplitSnailNumber("[[2,3],2]"), []string{"[2,3]", "2"})
	is.Equal(SplitSnailNumber("[[2,3],[4,5]]"), []string{"[2,3]", "[4,5]"})
	is.Equal(SplitSnailNumber("[[2,[6,7]],[4,5]]"), []string{"[2,[6,7]]", "[4,5]"})

	is.Equal(SplitSnailNumber("1"), []string{"1"})
}

func TestParseSnailNumber(t *testing.T) {
	is := is_.New(t)

	got := ParseSnailNumber("[1,2]")
	want := &SnailNumber{
		First:  &SnailNumber{Regular: 1},
		Second: &SnailNumber{Regular: 2},
	}

	if !cmp.Equal(want, got, ignoreParent) {
		t.Error(cmp.Diff(want, got, ignoreParent))
	}

	got = ParseSnailNumber("[1,[2,3]]")
	want = &SnailNumber{
		First: &SnailNumber{Regular: 1},
		Second: &SnailNumber{
			First:  &SnailNumber{Regular: 2},
			Second: &SnailNumber{Regular: 3},
		},
	}

	if !cmp.Equal(want, got, ignoreParent) {
		t.Error(cmp.Diff(want, got, ignoreParent))
	}

	is.Equal(got.First.Parent, got) // snail number has correct parent
}

func TestTraverseSnailLeft(t *testing.T) {
	is := is_.New(t)

	sn := ParseSnailNumber("[1,[2,3]]")
	left := TraverseSnailLeft(sn.Second.First)
	is.Equal(sn.First, left) // should find 1 on left

	sn = ParseSnailNumber("[[[1,2],3],4]")
	left = TraverseSnailLeft(sn.First.First.First)
	is.Equal(nil, left) // should find nil

	sn = ParseSnailNumber("[[[1,2],[[3,4],5]],6]")
	left = TraverseSnailLeft(sn.First.Second.First.First)
	is.Equal(sn.First.First.Second, left) // should 2 on left

}

func TestTraverseSnailRight(t *testing.T) {
	is := is_.New(t)

	sn := ParseSnailNumber("[[1,2],3]")
	right := TraverseSnailRight(sn.First.Second)
	is.Equal(sn.Second, right) // should find 1 on left

	sn = ParseSnailNumber("[1,[2,[3,4]]]")
	right = TraverseSnailRight(sn.Second.Second.Second)
	is.Equal(nil, right) // should find nil

	sn = ParseSnailNumber("[[[1,2],[[3,4],5]],6]")
	right = TraverseSnailRight(sn.First.First.Second)
	is.Equal(sn.First.Second.First.First, right) // should 2 on left

}

func TestSnailNumberDepth(t *testing.T) {
	is := is_.New(t)

	sn := ParseSnailNumber("[[1,2],3]")
	is.Equal(sn.First.Depth(), 1)

	sn = ParseSnailNumber("[[[[[9,8],1],2],3],4]")
	target := sn.First.First.First.First
	is.Equal(target.First.Regular, 9)
	is.Equal(target.Depth(), 4)
}

func TestSnailNumberExplode(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "[[[[[9,8],1],2],3],4]",
			out: "[[[[0,9],2],3],4]",
		},
		{
			in:  "[7,[6,[5,[4,[3,2]]]]]",
			out: "[7,[6,[5,[7,0]]]]",
		},
		{
			in:  "[[6,[5,[4,[3,2]]]],1]",
			out: "[[6,[5,[7,0]]],3]",
		},
		{
			in:  "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			out: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			in:  "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			out: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s=>%s", tt.in, tt.out), func(t *testing.T) {
			is := is_.New(t)

			got := ParseSnailNumber("[[[[[9,8],1],2],3],4]")
			want := ParseSnailNumber("[[[[0,9],2],3],4]")

			exploded := got.Explode()
			is.True(exploded) // should explode
			if !cmp.Equal(want, got, ignoreParent) {
				t.Error(cmp.Diff(want, got, ignoreParent))
			}

		})
	}
}

func TestSnailNumberSplit(t *testing.T) {
	is := is_.New(t)

	got := &SnailNumber{Regular: 10}
	splited := got.Split()
	is.True(splited)

	want := ParseSnailNumber("[5,5]")
	if !cmp.Equal(want, got, ignoreParent) {
		t.Error(cmp.Diff(want, got, ignoreParent))
	}

	got = &SnailNumber{Regular: 11}
	got.Split()
	is.True(splited)

	want = ParseSnailNumber("[5,6]")
	if !cmp.Equal(want, got, ignoreParent) {
		t.Error(cmp.Diff(want, got, ignoreParent))
	}
}

func TestSnailNumberReduce(t *testing.T) {
	got := ParseSnailNumber("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
	want := ParseSnailNumber("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")

	got.Reduce()

	if !cmp.Equal(want, got, ignoreParent) {
		t.Error(cmp.Diff(want, got, ignoreParent))
	}
}

func TestSumSnailNumbers(t *testing.T) {
	tests := []struct {
		in  []string
		out string
	}{
		{
			in: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
			},
			out: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			in: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
			},
			out: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			in: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
				"[6,6]",
			},
			out: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			in: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			},
			out: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SumSnailNumbers(tt.in)
			want := ParseSnailNumber(tt.out)

			if !cmp.Equal(want, got, ignoreParent) {
				t.Error(cmp.Diff(want, got, ignoreParent))
			}
		})
	}
}

func TestSnailNumberMagnitude(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "[[1,2],[[3,4],5]]",
			out: 143,
		}, {
			in:  "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			out: 1384,
		}, {
			in:  "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			out: 445,
		}, {
			in:  "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			out: 791,
		}, {
			in:  "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			out: 1137,
		}, {
			in:  "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			out: 3488,
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			is := is_.New(t)

			got := ParseSnailNumber(tt.in).Magnitude()
			is.Equal(got, tt.out)
		})
	}
}

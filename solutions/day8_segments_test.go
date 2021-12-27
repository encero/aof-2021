package solutions

import (
	is_ "github.com/matryer/is"
	"testing"
)

func TestNewSevenSegment(t *testing.T) {
	is := is_.New(t)

	segment := NewSevenSegment("be")

	is.Equal(segment, SB|SE)
}

func TestParseSegmentsLine(t *testing.T) {
	is := is_.New(t)

	segments := ParseSegmentsLine("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe")

	is.Equal(segments.numbers, []SevenSegment{
		NewSevenSegment("be"),
		NewSevenSegment("cfbegad"),
		NewSevenSegment("cbdgef"),
		NewSevenSegment("fgaecd"),
		NewSevenSegment("cgeb"),
		NewSevenSegment("fdcge"),
		NewSevenSegment("agebfd"),
		NewSevenSegment("fecdb"),
		NewSevenSegment("fabcd"),
		NewSevenSegment("edb"),
	})

	is.Equal(segments.result, []SevenSegment{
		NewSevenSegment("fdgacbe"),
		NewSevenSegment("cefdb"),
		NewSevenSegment("cefbgd"),
		NewSevenSegment("gcbe"),
	})

}

var segmentsInput = []SegmentsLine{
	ParseSegmentsLine("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"),
	ParseSegmentsLine("edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"),
	ParseSegmentsLine("fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg"),
	ParseSegmentsLine("fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb"),
	ParseSegmentsLine("aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea"),
	ParseSegmentsLine("fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb"),
	ParseSegmentsLine("dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe"),
	ParseSegmentsLine("bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef"),
	ParseSegmentsLine("egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb"),
	ParseSegmentsLine("gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"),
}

func TestSimpleSolve(t *testing.T) {
	is := is_.New(t)
	solutions := SolveAllSegments(segmentsInput, SimpleSolveSegments)

	is.Equal(CountSolvedSegments(solutions), 26)
}

func TestFullSolve(t *testing.T) {
	is := is_.New(t)
	solutions := SolveAllSegments(segmentsInput, FullSolveSegments)

	answer := SevenSegmentFinalAnswer(solutions)

	is.Equal(answer, 61229)

}

package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"strings"
)

func Day8SevenSegments() error {
	var input = []SegmentsLine{}
	helpers.ReadLines("inputs/day8.txt", func(line string) error {
		input = append(input, ParseSegmentsLine(line))

		return nil
	})

	simpleSolution := SolveAllSegments(input, SimpleSolveSegments)
	simpleNumberCount := CountSolvedSegments(simpleSolution)

	fmt.Println("Day 8 - Part 1:", simpleNumberCount)

	fullSolution := SolveAllSegments(input, FullSolveSegments)
	answer := SevenSegmentFinalAnswer(fullSolution)

	fmt.Println("Day 8 - Part 2:", answer)

	return nil
}

const (
	SA SevenSegment = 1 << iota
	SB
	SC
	SD
	SE
	SF
	SG
)

func SevenSegmentFinalAnswer(solutions []SegmentsSolution) int {
	var answer int
	for _, solution := range solutions {
		answer += solution.ResultNumber()
	}

	return answer
}

type SevenSegment int

func (s SevenSegment) String() string {
	out := make([]string, 0, 7)
	if s&SA > 0 {
		out = append(out, "A")
	}
	if s&SB > 0 {
		out = append(out, "B")
	}
	if s&SC > 0 {
		out = append(out, "C")
	}
	if s&SD > 0 {
		out = append(out, "D")
	}
	if s&SE > 0 {
		out = append(out, "E")
	}
	if s&SF > 0 {
		out = append(out, "F")
	}
	if s&SG > 0 {
		out = append(out, "G")
	}

	return strings.Join(out, "")
}

func NewSevenSegment(in string) SevenSegment {
	var segment SevenSegment
	for _, c := range strings.Split(in, "") {
		switch strings.ToUpper(c) {
		case "A":
			segment |= SA
		case "B":
			segment |= SB
		case "C":
			segment |= SC
		case "D":
			segment |= SD
		case "E":
			segment |= SE
		case "F":
			segment |= SF
		case "G":
			segment |= SG
		}
	}

	return segment
}

type SegmentsLine struct {
	numbers []SevenSegment
	result  []SevenSegment
}

func ParseSegmentsLine(line string) SegmentsLine {
	var segments SegmentsLine

	parts := strings.Split(line, "|")

	numbers := helpers.FilterEmptyStrings(strings.Split(parts[0], " "))
	for _, number := range numbers {
		segments.numbers = append(segments.numbers, NewSevenSegment(number))
	}

	results := helpers.FilterEmptyStrings(strings.Split(parts[1], " "))

	for _, result := range results {
		segments.result = append(segments.result, NewSevenSegment(result))
	}

	return segments
}

type SegmentsSolution struct {
	solvedNumbers map[int]SevenSegment

	numbers []int
	result  []int
}

func SolveAllSegments(lines []SegmentsLine, solver func(line SegmentsLine) SegmentsSolution) []SegmentsSolution {
	var solutions []SegmentsSolution

	for _, line := range lines {
		solution := solver(line)
		solutions = append(solutions, solution)
	}

	return solutions
}

func CountSolvedSegments(solutions []SegmentsSolution) int {
	var count int

	for _, solution := range solutions {
		for _, n := range solution.result {
			if n != -1 {
				count++
			}
		}
	}

	return count
}

func SimpleSolveSegments(line SegmentsLine) SegmentsSolution {
	solution := SegmentsSolution{
		solvedNumbers: make(map[int]SevenSegment),
		numbers:       make([]int, len(line.numbers)),
		result:        make([]int, len(line.result)),
	}

	solution.AssignSimpleNumbers(line.numbers)

	solution.result = solution.decode(line.result)

	return solution
}

func FullSolveSegments(line SegmentsLine) SegmentsSolution {
	solution := SegmentsSolution{
		solvedNumbers: make(map[int]SevenSegment),
		numbers:       make([]int, len(line.numbers)),
		result:        make([]int, len(line.result)),
	}

	solution.AssignSimpleNumbers(line.numbers)

	solution.SolveForThree(line.numbers)
	solution.SolveForNine(line.numbers)
	solution.SolveForSix(line.numbers)
	solution.SolveForTwo(line.numbers)
	solution.SolveForFive(line.numbers)
	solution.SolveForZero(line.numbers)

	solution.result = solution.decode(line.result)

	return solution
}

func (s SegmentsSolution) ResultNumber() int {
	return s.result[0]*1000 + s.result[1]*100 + s.result[2]*10 + s.result[3]
}

func (s SegmentsSolution) nSegmentNumbers(digits []SevenSegment, segments int, f func(digit SevenSegment)) {
	for _, digit := range digits {
		if len(digit.String()) == segments {
			f(digit)
		}
	}
}

func (s *SegmentsSolution) SolveForFive(digits []SevenSegment) {
	s.nSegmentNumbers(digits, 5, func(digit SevenSegment) {
		if digit != s.solvedNumbers[2] && digit != s.solvedNumbers[3] {
			s.solvedNumbers[5] = digit
		}
	})
}

func (s *SegmentsSolution) SolveForTwo(digits []SevenSegment) {
	// real segment C is difference between 6 and eight
	segmentC := s.solvedNumbers[8] ^ s.solvedNumbers[6]&0b1111111

	s.nSegmentNumbers(digits, 5, func(digit SevenSegment) {
		if digit&segmentC == segmentC && digit != s.solvedNumbers[3] {
			s.solvedNumbers[2] = digit
		}
	})
}

func (s *SegmentsSolution) SolveForZero(digits []SevenSegment) {
	s.nSegmentNumbers(digits, 6, func(digit SevenSegment) {
		// zero is six segment digit that isn't six or nine
		if digit != s.solvedNumbers[9] && digit != s.solvedNumbers[6] {
			s.solvedNumbers[0] = digit
		}
	})
}

func (s *SegmentsSolution) SolveForSix(digits []SevenSegment) {
	s.nSegmentNumbers(digits, 6, func(digit SevenSegment) {
		// six is only six segment digit that doesn't contain digit 1
		if digit&s.solvedNumbers[1] != s.solvedNumbers[1] {
			s.solvedNumbers[6] = digit
		}
	})
}

func (s *SegmentsSolution) SolveForNine(digits []SevenSegment) {
	s.nSegmentNumbers(digits, 6, func(digit SevenSegment) {
		// nine is only six segment digit that contains digit 4
		if digit&s.solvedNumbers[4] == s.solvedNumbers[4] {
			s.solvedNumbers[9] = digit
		}
	})
}

func (s *SegmentsSolution) SolveForThree(digits []SevenSegment) {
	s.nSegmentNumbers(digits, 5, func(digit SevenSegment) {
		// three is only five segment digit that contains digit 1
		if digit&s.solvedNumbers[1] == s.solvedNumbers[1] {
			s.solvedNumbers[3] = digit
		}
	})
}

func (s *SegmentsSolution) AssignSimpleNumbers(digits []SevenSegment) {
	for _, n := range digits {
		switch len(n.String()) {
		case 2:
			s.solvedNumbers[1] = n
		case 3:
			s.solvedNumbers[7] = n
		case 4:
			s.solvedNumbers[4] = n
		case 7:
			s.solvedNumbers[8] = n
		}
	}
}

func (s SegmentsSolution) decode(result []SevenSegment) []int {
	var out []int

	for _, segment := range result {
		for decoded, digit := range s.solvedNumbers {
			if segment == digit {
				out = append(out, decoded)
			}
		}
	}

	return out
}

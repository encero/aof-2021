package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"sort"
	"strings"
)


func Day10SyntaxScoring() error {
	var input []string
	helpers.ReadLines("inputs/day10.txt", func(s string) error {
		input = append(input, s)
		return nil
	})

	score := SyntaxCheckPartOne(input)
	fmt.Println("Day 10, part 1:", score)

	score = SyntaxCheckPartTwo(input)
	fmt.Println("Day 10, part 2:", score)

	return nil
}

type SyntaxError struct {
	Type string
	Char string
}

var ErrorTypeCorrupt = "corrupt"

func SyntaxCheck(input string) ([]string, *SyntaxError) {
	stackPointer := 0
	stack := make([]string, len(input))

	chars := strings.Split(input, "")

	for _, v := range chars {
		if isOpenBracket(v) {
			stack[stackPointer] = v
			stackPointer++
			continue
		}

		closing := bracketPair(stack[stackPointer-1])
		if v == closing {
			stackPointer--
			continue
		}

		return nil, &SyntaxError{
			Type: ErrorTypeCorrupt,
			Char: v,
		}
	}

	return stack[:stackPointer], nil
}

func SyntaxCheckPartOne(input []string) int {
	var score int

	for _, line := range input {
		if _, err := SyntaxCheck(line); err != nil && err.Type == ErrorTypeCorrupt {
			score += corruptBracketScore(err.Char)
		}
	}

	return score
}

func SyntaxCheckPartTwo(input []string) int {
	scoreF := func(stack []string) int {
		score := 0

		for i := len(stack) - 1; i >= 0; i-- {
			score = (score * 5) + missingBracketScore(bracketPair(stack[i]))
		}

		return score
	}

	var scores []int

	for _, line := range input {
		stack, err := SyntaxCheck(line)
		if err != nil {
			continue
		}

		scores = append(scores, scoreF(stack))
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func isOpenBracket(char string) bool {
	return char == "(" || char == "[" || char == "{" || char == "<"
}

func bracketPair(char string) string {
	switch char {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	}

	panic("Invalid bracket")
}

func corruptBracketScore(char string) int {
	switch char {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}

	panic("Invalid bracket")
}

func missingBracketScore(char string) int {
	switch char {
	case ")":
		return 1
	case "]":
		return 2
	case "}":
		return 3
	case ">":
		return 4
	}

	panic("Invalid bracket")
}

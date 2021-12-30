package solutions

import (
	is_ "github.com/matryer/is"
	"testing"
)

func TestSyntaxCheck(t *testing.T) {
	is := is_.New(t)

	_, err := SyntaxCheck("{([(<{}[<>[]}>{[]{[(<()>")

	is.Equal(err.Type, ErrorTypeCorrupt)
	is.Equal(err.Char, "}")

	_, err = SyntaxCheck("[[<[([]))<([[{}[[()]]]")

	is.Equal(err.Type, ErrorTypeCorrupt)
	is.Equal(err.Char, ")")

	_, err = SyntaxCheck("[{[{({}]{}}([{[{{{}}([]")

	is.Equal(err.Type, ErrorTypeCorrupt)
	is.Equal(err.Char, "]")

	_, err = SyntaxCheck("[<(<(<(<{}))><([]([]()")

	is.Equal(err.Type, ErrorTypeCorrupt)
	is.Equal(err.Char, ")")
}

var syntaxTestInput = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

func TestSyntaxCheckPartOne(t *testing.T) {
	is := is_.New(t)

	score := SyntaxCheckPartOne(syntaxTestInput)

	is.Equal(score, 26397)
}

func TestSyntaxCheckPartTwo(t *testing.T) {
	is := is_.New(t)

	score := SyntaxCheckPartTwo(syntaxTestInput)

	is.Equal(score, 288957)
}

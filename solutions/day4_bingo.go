package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"io/ioutil"
	"strings"
)

func loadBoards() []*Board {
	prefix := "../"
	content, err := ioutil.ReadFile(prefix + "inputs/day4_boards.txt")
	if err != nil {
		panic(err)
	}

	data := helpers.StringsToInts(
		helpers.FilterEmptyStrings(
			strings.Split(
				strings.ReplaceAll(
					strings.ReplaceAll(string(content), "\n", " "),
					"  ", " ",
				),
				" ",
			),
		),
	)

	boards := make([]*Board, 0, len(data)/25)
	for i := 0; i < len(data); i += 25 {
		boards = append(boards, NewBoard(data[i:i+25]))
	}

	return boards
}

func Day4Bingo() error {
	numbers := helpers.StringsToInts(strings.Split(helpers.ReadLine("inputs/day4_numbers.txt"), ","))
	boards := loadBoards()

	_, board, called := FindFirstWinningBingoBoard(numbers, boards)

	sum := board.SumNotCrossedNumbers()
	fmt.Println("Day 4 part 1: ", sum*called)

	boards = loadBoards()
	board, called = FindLastWinningBoard(numbers, boards)
	fmt.Println("Day 4 part 2: ", board.SumNotCrossedNumbers()*called)

	return nil
}

func FindFirstWinningBingoBoard(numbers []int, boards []*Board) (int, *Board, int) {
	for _, number := range numbers {
		for i, board := range boards {
			win := board.Call(number)
			if win {
				return i, board, number
			}
		}
	}

	panic("un solvable")
}

func FindLastWinningBoard(numbers []int, boards []*Board) (*Board, int) {
	filter := func(boards []*Board, number int) []*Board {
		notWinning := make([]*Board, 0, len(boards))
		for _, board := range boards {
			if !board.Call(number) {
				notWinning = append(notWinning, board)
			}
		}

		return notWinning
	}

	for _, number := range numbers {
		if len(boards) != 1 {
			boards = filter(boards, number)
			continue
		}

		if boards[0].Call(number) {
			return boards[0], number
		}
	}

	panic("un solvable")
}

type Board struct {
	numbers    []int
	notCrossed [][]int
	crossed    [][]int
}

func NewBoard(numbers []int) *Board {
	var notCrossed [][]int

	for row := 0; row < 5; row++ {
		notCrossed = append(notCrossed, append([]int{}, numbers[row*5:(row+1)*5]...))
	}

	for col := 0; col < 5; col++ {
		notCrossed = append(notCrossed, append([]int{},
			numbers[col],
			numbers[col+5],
			numbers[col+10],
			numbers[col+15],
			numbers[col+20],
		))
	}

	return &Board{
		numbers:    numbers,
		crossed:    make([][]int, 5*5),
		notCrossed: notCrossed,
	}
}

func (b *Board) Call(number int) bool {
	win := false

	for i := range b.notCrossed {
		if b.cross(i, number) {
			win = win || len(b.notCrossed[i]) == 0
		}
	}

	return win
}

func (b Board) cross(i, number int) bool {
	index, ok := helpers.IndexOfInt(b.notCrossed[i], number)
	if !ok {
		return false
	}

	b.notCrossed[i] = append(b.notCrossed[i][:index], b.notCrossed[i][index+1:]...)
	b.crossed[i] = append(b.crossed[i], number)

	return true
}

func (b *Board) SumNotCrossedNumbers() int {
	uniqueNotCrossed := make(map[int]struct{})
	for _, numbers := range b.notCrossed {
		for _, number := range numbers {
			uniqueNotCrossed[number] = struct{}{}
		}
	}

	var sum int
	for number := range uniqueNotCrossed {
		sum += number
	}

	return sum
}

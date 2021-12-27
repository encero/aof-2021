package solutions

import (
	is_ "github.com/matryer/is"
	"testing"
)

func TestBingoBoard(t *testing.T) {
	is := is_.New(t)

	board := NewBoard([]int{
		36, 11, 70, 77, 80,
		63, 3, 56, 75, 28,
		89, 91, 27, 33, 82,
		53, 79, 52, 96, 32,
		58, 14, 78, 65, 38,
	})

	is.True(!board.Call(36))
	is.True(!board.Call(11))
	is.True(!board.Call(70))
	is.True(!board.Call(77))
	is.True(board.Call(80))
}

func TestLoadBoards(t *testing.T) {
	is := is_.New(t)

	boards := loadBoards()

	is.True(len(boards) > 0)

	is.Equal(boards[0].numbers, []int{
		36, 11, 70, 77, 80,
		63, 3, 56, 75, 28,
		89, 91, 27, 33, 82,
		53, 79, 52, 96, 32,
		58, 14, 78, 65, 38,
	})
}

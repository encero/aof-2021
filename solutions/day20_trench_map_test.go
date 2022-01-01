package solutions

import (
	"testing"

	"github.com/encero/advent-of-code-2021/helpers"
	is_ "github.com/matryer/is"
)


func TestBoundsFromPixelMap(t *testing.T) {
    is := is_.New(t)

    bounds := BoundsFromPixelMap(map[helpers.Vec2]struct{}{
        {X: 5, Y:10}: {},
        {X: 7, Y:3}: {},
    })

    is.Equal(bounds, Bounds{X0: 5, X1:7, Y0: 3, Y1:10})
}

func TestImageEnhance(t *testing.T){
    is := is_.New(t)

    algo := ImageAlgoFromString("..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#")
    is.Equal(len(algo), 512) // algo char count

    pixels := ImagePixelsFromString(`#..#.
#....
##..#
..#..
..###`)
    is.Equal(len(pixels), 10) // image lit pixels

    im := NewImage(pixels)

    im.Enhance(algo)
    im.Enhance(algo)

    is.Equal(len(im.Pixels), 35) // enhanced lit pixels

    for i := 2; i < 50; i++ {
        im.Enhance(algo)
    }

    is.Equal(len(im.Pixels), 3351)
}


package solutions

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/encero/advent-of-code-2021/helpers"
)

func Day20TrenchMap() error {
	algo := ImageAlgoFromString(helpers.ReadLine("inputs/day20_algo.txt"))

	data, err := os.ReadFile("inputs/day20_image.txt")
	if err != nil {
		return err
	}
	pixels := ImagePixelsFromString(string(data))

	im := NewImage(pixels)

	fmt.Println("loaded", len(im.Pixels))
	im.Enhance(algo)
	fmt.Println("first enhance", len(im.Pixels))
	im.Enhance(algo)
	fmt.Println("second enhance", len(im.Pixels))

	fmt.Println("lit image pixels:", len(im.Pixels))

	for i := 2; i < 50; i++ {
		im.Enhance(algo)
	}

	fmt.Println("lit image after 50 enhances:", len(im.Pixels))

	return nil
}

type Image struct {
	Pixels     map[helpers.Vec2]struct{}
	Bounds     Bounds
	Background bool
}

func NewImage(pixels map[helpers.Vec2]struct{}) *Image {
	return &Image{
		Pixels: pixels,
		Bounds: BoundsFromPixelMap(pixels),
	}
}

func (i *Image) Enhance(algo []bool) {

	oldBounds := i.Bounds

	i.Bounds.X0--
	i.Bounds.X1++

	i.Bounds.Y0--
	i.Bounds.Y1++

	out := make(map[helpers.Vec2]struct{})

	for x := i.Bounds.X0; x <= i.Bounds.X1; x++ {
		for y := i.Bounds.Y0; y <= i.Bounds.Y1; y++ {
			target := helpers.Vec2{X: x, Y: y}
			index := i.AlgoIndex(target, oldBounds)

			if algo[index] {
				out[target] = struct{}{}
			}
		}
	}

	i.Pixels = out

	if algo[0] {
		i.Background = !i.Background
	}
}

func (im *Image) AlgoIndex(at helpers.Vec2, bounds Bounds) int {
	out := 0

	for i, v := range helpers.NineDirections {
		pix := at.Add(v)
		fakeBackground := im.Background && !bounds.TrueHit(pix)

		if _, ok := im.Pixels[pix]; ok || fakeBackground {
			out |= 1 << (8 - i)
		}
	}

	return out
}

func BoundsFromPixelMap(pixels map[helpers.Vec2]struct{}) Bounds {
	once := sync.Once{}

	var bounds Bounds

	for v := range pixels {
		once.Do(func() {
			bounds = Bounds{
				X0: v.X,
				X1: v.X,
				Y0: v.Y,
				Y1: v.Y,
			}
		})

		if v.X < bounds.X0 {
			bounds.X0 = v.X
		}
		if v.X > bounds.X1 {
			bounds.X1 = v.X
		}
		if v.Y < bounds.Y0 {
			bounds.Y0 = v.Y
		}
		if v.Y > bounds.Y1 {
			bounds.Y1 = v.Y
		}

	}

	return bounds
}

func ImageAlgoFromString(in string) []bool {
	algo := make([]bool, 512)

	for i, v := range in {
		if string(v) == "#" {
			algo[i] = true
		}
	}

	return algo
}

func ImagePixelsFromString(in string) map[helpers.Vec2]struct{} {
	pixMap := make(map[helpers.Vec2]struct{})

	lines := strings.Split(in, "\n")
	for y, line := range lines {
		for x, ch := range line {
			if string(ch) == "#" {
				pixMap[helpers.Vec2{X: x, Y: y}] = struct{}{}
			}
		}
	}

	return pixMap
}

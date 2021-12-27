package solutions

import (
	is_ "github.com/matryer/is"
	"testing"
)

func TestSliceDistance(t *testing.T) {
	is := is_.New(t)

	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	simpleSum := func(i int) int {
		return i
	}



	is.Equal(SliceDistance(2, input, simpleSum), 37)
	is.Equal(SliceDistance(1, input, simpleSum), 41)


	is.Equal(SliceDistance(5, input, CrabDistance), 168)
	is.Equal(SliceDistance(2, input, CrabDistance), 206)
}

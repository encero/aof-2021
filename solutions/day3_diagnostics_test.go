package solutions

import (
	is_ "github.com/matryer/is"
	"testing"
)

func TestAverageBitValues(t *testing.T) {
	is := is_.New(t)

	input := []int{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}

	gamma, epsilon := AverageBitValues(input, 0b11111)

	is.Equal(gamma, 0b10110)   // gamma
	is.Equal(epsilon, 0b01001) // epsilon
}

func TestDiagnosticsBitsFilter(t *testing.T) {
	is := is_.New(t)

	input := []int{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}

	oxy := DiagnosticsBitsFilter(input, func(ints []int) int {
		gamma, _ := AverageBitValues(ints, 0b11111)
		return gamma
	})
	is.Equal(oxy, 0b10111) // oxy

	co2 := DiagnosticsBitsFilter(input, func(ints []int) int {
		_, epsilon := AverageBitValues(ints, 0b11111)
		return epsilon
	})
	is.Equal(co2, 0b01010) // co2
}

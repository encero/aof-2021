package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"strconv"
)

const DiagnosticsMask = 0b111111111111

func Day3Diagnostics() error {
	var input []int
	helpers.ReadLines("inputs/day3.txt", func(s string) error {
		i, err := strconv.ParseInt(s, 2, 32)
		if err != nil {
			return fmt.Errorf("day3 input parse error: %w", err)
		}

		input = append(input, int(i))

		return nil
	})

	gamma, epsilon := AverageBitValues(input, DiagnosticsMask)

	fmt.Println("Day 3, part 1:", gamma*epsilon)

	oxy := DiagnosticsBitsFilter(input, func(ints []int) int {
		gamma, _ := AverageBitValues(ints, DiagnosticsMask)
		return gamma
	})

	co2 := DiagnosticsBitsFilter(input, func(ints []int) int {
		_, epsilon := AverageBitValues(ints, DiagnosticsMask)
		return epsilon
	})

	fmt.Println("Day 3, part 2:", oxy*co2)

	return nil
}

func AverageBitValues(input []int, mask int) (int, int) {
	countOnes := make([]int, 32)

	for _, v := range input {
		for i := 0; i < 32; i++ {
			if (v & (1 << uint(i))) != 0 {
				countOnes[i]++
			}
		}
	}

	return Gama(countOnes, len(input)), ^Gama(countOnes, len(input)) & mask
}

func Gama(input []int, count int) int {
	var out int

	for i, v := range input {
		if v >= count-v {
			out |= 1 << uint(i)
		}
	}

	return out
}

func DiagnosticsBitsFilter(input []int, averager func([]int) int) int {
	filter := func(input []int, target, mask int) []int {
		var out []int

		for _, v := range input {
			masked := v & mask

			if masked == target {
				out = append(out, v)
			}
		}

		return out
	}

	var mask int
	for i := 0; i < 32; i++ {
		target := averager(input)

		mask = 1 << (31 - i)

		input = filter(input, mask&target, mask)
		if len(input) == 1 {
			return input[0]
		}
	}

	panic("no match found")
}

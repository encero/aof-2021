package solutions

import (
	"fmt"
	"strings"

	"github.com/encero/advent-of-code-2021/helpers"
)

func Day14Polymers() error {
	pairs := make(map[string]string)

	helpers.ReadLines("inputs/day14.txt", func(line string) error {
		if line == "" {
			return nil
		}

		splits := strings.Split(line, " -> ")

		pairs[splits[0]] = splits[1]

		return nil
	})

	polymer := PolymerToPairs("SVKVKCCBNHNSOSCCOPOC")

	for i := 0; i < 40; i++ {
		polymer = Polymerize2(polymer, pairs)
	}

	PolymerMinMax(polymer)

	return nil
}

func PolymerMinMax(polymer map[string]int64) int64 {

	elements := make(map[string]int64)

	for p, count := range polymer {
		elements[string(p[0])] += count / 2
		elements[string(p[1])] += count / 2
	}

	min := int64(9223372036854775807)
	max := int64(0)

	for _, v := range elements {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max - min
}

func PolymerToPairs(polymer string) map[string]int64 {
	pairs := make(map[string]int64)

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		pairs[pair] += 1
	}

	return pairs
}

func Polymerize2(polymer map[string]int64, pairs map[string]string) map[string]int64 {
	newPoly := make(map[string]int64)

	for p, count := range polymer {
		if insert, ok := pairs[p]; ok {
			first := fmt.Sprintf("%s%s", string(p[0]), insert)
			second := fmt.Sprintf("%s%s", insert, string(p[1]))

			newPoly[first] += count
			newPoly[second] += count
			polymer[p] = 0
		}
	}

	for p, count := range polymer {
		if count > 0 {
			newPoly[p] += count
		}
	}

	return newPoly
}

func Polymerize(polymer string, pairs map[string]string) string {
	chars := strings.Split(polymer, "")

	for i := 0; i < len(chars)-1; i++ {
		pair := strings.Join(chars[i:i+2], "")

		insertion, ok := pairs[pair]
		if ok {
			chars = append(chars[:i+1], append([]string{insertion}, chars[i+1:]...)...)
			i++
		}
	}

	return strings.Join(chars, "")
}

func CountPolymers(polymer string) map[string]int {
	counts := make(map[string]int)

	for _, v := range strings.Split(polymer, "") {
		if _, ok := counts[v]; !ok {
			counts[v] = 0
		}

		counts[v] = counts[v] + 1
	}

	return counts
}

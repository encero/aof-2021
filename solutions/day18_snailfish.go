package solutions

import (
	"fmt"
	"strconv"

	"github.com/encero/advent-of-code-2021/helpers"
)

func Day18SnailNumbers() error {
	nums, err := helpers.ReadAllLines("inputs/day18.txt")
	if err != nil {
		return nil
	}

	sum := SumSnailNumbers(nums)

	fmt.Println("Day 18, Part 1:", sum.Magnitude())

	max := 0

	for _, a := range nums {
		for _, b := range nums {
			if a == b {
				continue
			}

			if mg := SumSnailNumbers([]string{a, b}).Magnitude(); mg > max {
				max = mg
			}

			if mg := SumSnailNumbers([]string{b, a}).Magnitude(); mg > max {
				max = mg
			}
		}
	}

	fmt.Println("Part 2:", max)

	return nil
}

type SnailNumber struct {
	First   *SnailNumber
	Second  *SnailNumber
	Regular int
	Parent  *SnailNumber
}

func (s SnailNumber) IsRegular() bool {
	return s.First == nil && s.Second == nil
}

func (s SnailNumber) Depth() int {
	depth := 0
	for s.Parent != nil {
		depth++
		s = *s.Parent
	}

	return depth
}

func (s *SnailNumber) Explode() bool {
	if s.IsRegular() {
		return false
	}

	if s.Depth() == 4 {
		left := TraverseSnailLeft(s)
		right := TraverseSnailRight(s)

		if left != nil {
			left.Regular += s.First.Regular
		}

		if right != nil {
			right.Regular += s.Second.Regular
		}

		s.First = nil
		s.Second = nil

		return true
	}
	if s.First.Explode() {
		return true
	}
	if s.Second.Explode() {
		return true
	}

	return false
}

func (s *SnailNumber) Split() bool {
	if s.IsRegular() && s.Regular >= 10 {
		s.First = &SnailNumber{Regular: s.Regular / 2, Parent: s}
		s.Second = &SnailNumber{Regular: s.Regular/2 + s.Regular%2, Parent: s}
		s.Regular = 0

		return true
	}

	if s.First != nil && s.First.Split() {
		return true
	}

	if s.Second != nil && s.Second.Split() {
		return true
	}

	return false
}

func (s *SnailNumber) Reduce() {
	for {
		if s.Explode() {
			continue
		}

		if s.Split() {
			continue
		}

		break
	}
}

func (s *SnailNumber) Magnitude() int {
	if s.IsRegular() {
		return s.Regular
	}

	return s.First.Magnitude()*3 + s.Second.Magnitude()*2
}

func SumSnailNumbers(nums []string) *SnailNumber {
	sum := ParseSnailNumber(nums[0])

	for i := 1; i < len(nums); i++ {
		tmp := &SnailNumber{
			First:  sum,
			Second: ParseSnailNumber(nums[i]),
		}

		tmp.First.Parent = tmp
		tmp.Second.Parent = tmp

		tmp.Reduce()

		sum = tmp
	}

	return sum
}

func ParseSnailNumber(num string) *SnailNumber {
	splits := SplitSnailNumber(num)

	if len(splits) == 1 {
		num, _ := strconv.Atoi(num)
		return &SnailNumber{Regular: num}
	}

	first := ParseSnailNumber(splits[0])
	second := ParseSnailNumber(splits[1])
	snail := &SnailNumber{
		First:  first,
		Second: second,
	}

	first.Parent = snail
	second.Parent = snail

	return snail
}

func SplitSnailNumber(s string) []string {
	if string(s[0]) != "[" {
		return []string{s}
	}

	out := make([]string, 0, 2)

	depth := 0
	comma := 0

loop:
	for i := 1; i < len(s); i++ {
		ch := string(s[i])

		switch ch {
		case "[":
			depth++
		case "]":
			depth--
		case ",":
			if depth == 0 {
				comma = i
				break loop
			}
		}
	}

	out = append(out, s[1:comma])
	out = append(out, s[comma+1:len(s)-1])

	return out
}

func TraverseSnailLeft(snail *SnailNumber) *SnailNumber {
	prev := snail
	//if snail.IsRegular() {
	//	fmt.Println("escaping regular")
	//	snail = snail.Parent
	//}

	// traverse up the stack
	for snail.Parent != nil {
		prev = snail
		snail = snail.Parent
		if snail.Second == prev {
			break
		}
	}

	// no regular found until top
	if snail.Parent == nil && snail.First == prev {
		return nil
	}

	// top pair has regular on left a we came from right
	if snail.Parent == nil && snail.First.IsRegular() {
		return snail.First
	}

	// start search in left stack
	snail = snail.First

	// traverse pairs down "right" until first regular
	for !snail.IsRegular() {
		snail = snail.Second
	}

	return snail
}

func TraverseSnailRight(snail *SnailNumber) *SnailNumber {
	prev := snail

	// traverse up the stack
	for snail.Parent != nil {
		prev = snail
		snail = snail.Parent
		if snail.First == prev {
			break
		}
	}

	// no regular found until top
	if snail.Parent == nil && snail.Second == prev {
		return nil
	}

	// top pair has regular on left a we came from right
	if snail.Parent == nil && snail.Second.IsRegular() {
		return snail.Second
	}

	// start search in left stack
	snail = snail.Second

	// traverse pairs down "left" until first regular
	for !snail.IsRegular() {
		snail = snail.First
	}

	return snail
}

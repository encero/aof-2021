package main

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/solutions"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Advent of code 2021")

	if len(os.Args) == 1 {
		fmt.Printf("usage %s day\n", os.Args[0])
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("day is not a number\n")
		return
	}

	solution, ok := solutions.Solutions[day]
	if !ok {
		fmt.Printf("solution for day %d not found\n", day)
		return
	}

	err = solution()
	if err != nil {
		fmt.Printf("solution error: %v\n", err)
	}
}

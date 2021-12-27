package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLine(name string) string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func ReadLines(name string, fn func(string) error) error {
	inputFile, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("input file err: %w", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		err := fn(scanner.Text())
		if err != nil {
			return err
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

func ReadAllLines(name string) ([]string, error) {
	var lines []string

	err := ReadLines(name, func(line string) error {
		lines = append(lines, line)
		return nil
	})

	return lines, err
}

func ReadIntGrid(name string) [][]int {
	var input [][]int
	ReadLines(name, func(s string) error {
		row := StringsToInts(strings.Split(s, ""))
		input = append(input, row)

		return nil
	})

	return input
}

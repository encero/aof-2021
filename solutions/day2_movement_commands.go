package solutions

import (
	"fmt"
	"github.com/encero/advent-of-code-2021/helpers"
	"strconv"
	"strings"
)

type MovementDirection int

const (
	Up MovementDirection = iota
	Down
	Forward
)

type MovementCommand struct {
	Direction MovementDirection
	Distance  int
}

func Day2() error {
	var commands []MovementCommand

	err := helpers.ReadLines("inputs/day2.txt", func(s string) error {
		command, err := parseMovementCommand(s)
		if err != nil {
			return err
		}

		commands = append(commands, command)

		return nil
	})
	if err != nil {
		return err
	}

	depth, x := SimulateMovement(commands)
	fmt.Printf("Day 2 Part 1: %d\n", depth*x)

	depth, x = SimulateAimMovement(commands)
	fmt.Printf("Day 2 Part 2: %d\n", depth*x)

	return nil
}

func SimulateAimMovement(commands []MovementCommand) (int, int) {
	aim, depth, x := 0, 0, 0
	for _, command := range commands {
		switch command.Direction {
		case Up:
			aim -= command.Distance
		case Down:
			aim += command.Distance
		case Forward:
			x += command.Distance
			depth += aim * command.Distance
		}
	}

	return depth, x
}

func SimulateMovement(commands []MovementCommand) (int, int) {
	depth, x := 0, 0
	for _, command := range commands {
		switch command.Direction {
		case Up:
			depth -= command.Distance
		case Down:
			depth += command.Distance
		case Forward:
			x += command.Distance
		}
	}

	return depth, x
}

func parseDirection(s string) MovementDirection {
	switch s {
	case "up":
		return Up
	case "down":
		return Down
	case "forward":
		return Forward
	}

	panic("Unknown direction")
}

func parseMovementCommand(s string) (MovementCommand, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return MovementCommand{}, fmt.Errorf("invalid input: %s", s)
	}

	command := parseDirection(parts[0])
	distance, err := strconv.Atoi(parts[1])
	if err != nil {
		return MovementCommand{}, fmt.Errorf("invalid distance: %s", s)
	}

	return MovementCommand{
		Direction: command,
		Distance:  distance,
	}, nil
}

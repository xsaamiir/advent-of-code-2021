package day2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Day2 struct{}

type Direction int

const (
	DirectionUnknown Direction = iota
	DirectionForward
	DirectionDown
	DirectionUp
)

func NewDirectionFromString(s string) (Direction, error) {
	switch s {
	case "forward":
		return DirectionForward, nil
	case "down":
		return DirectionDown, nil
	case "up":
		return DirectionUp, nil
	default:
		return DirectionUnknown, fmt.Errorf("error parsing direction from string %s", s)
	}
}

type Step struct {
	Direction Direction
	Unit      uint
}

func (d Day2) InputGenerator(input io.Reader) (interface{}, error) {
	scanner := bufio.NewScanner(input)

	var output []Step

	for scanner.Scan() {
		l := scanner.Text()
		ss := strings.Split(l, " ")
		if len(ss) != 2 {
			return nil, fmt.Errorf("unexpted line format")
		}

		d, err := NewDirectionFromString(ss[0])
		if err != nil {
			return nil, err
		}

		u, err := strconv.Atoi(ss[1])
		if err != nil {
			return nil, err
		}

		output = append(output, Step{Direction: d, Unit: uint(u)})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return output, nil
}

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (p Position) Multiply() int {
	return p.Horizontal * p.Depth
}

func (d Day2) SolverPart1(v interface{}) (interface{}, error) {
	input := v.([]Step)

	var p Position

	for _, step := range input {
		switch step.Direction {
		case DirectionForward:
			p.Horizontal += int(step.Unit)
		case DirectionDown:
			p.Depth += int(step.Unit)
		case DirectionUp:
			p.Depth -= int(step.Unit)
		}
	}

	return p.Multiply(), nil
}

func (d Day2) SolverPart2(v interface{}) (interface{}, error) {
	input := v.([]Step)

	var p Position

	for _, step := range input {
		switch step.Direction {
		case DirectionForward:
			p.Horizontal += int(step.Unit)
			p.Depth += p.Aim * int(step.Unit)
		case DirectionDown:
			p.Aim += int(step.Unit)
		case DirectionUp:
			p.Aim -= int(step.Unit)
		}
	}

	return p.Multiply(), nil
}

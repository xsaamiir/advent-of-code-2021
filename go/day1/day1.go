package day1

import (
	"bufio"
	"io"
	"strconv"
)

type Day1 struct{}

func (d Day1) InputGenerator(input io.Reader) (interface{}, error) {
	scanner := bufio.NewScanner(input)

	var output []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		output = append(output, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return output, nil
}

func (d Day1) SolverPart1(v interface{}) (interface{}, error) {
	input := v.([]int)

	var c int

	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			c++
		}
	}

	return c, nil
}

func (d Day1) SolverPart2(v interface{}) (interface{}, error) {
	input := v.([]int)

	var s []int

	for i := 0; i < len(input)-2; i++ {
		s = append(s, input[i]+input[i+1]+input[i+2])
	}

	return d.SolverPart1(s)
}

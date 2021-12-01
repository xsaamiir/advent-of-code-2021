package aoc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func intputGeneratorDay1(input io.Reader) ([]int, error) {
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

func solverDay1Part1(input []int) int {
	var c int

	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			c++
		}
	}

	return c
}

func solverDay1Part2(input []int) int {
	var s []int

	for i := 0; i < len(input)-2; i++ {
		s = append(s, input[i]+input[i+1]+input[i+2])
	}

	return solverDay1Part1(s)
}

func mainDay1Part1() {
	f, err := os.Open("./input/day1.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	output, err := intputGeneratorDay1(f)
	if err != nil {
		panic(err)
	}

	a1 := solverDay1Part1(output)
	fmt.Println("Part 1 : ", a1)

	a2 := solverDay1Part2(output)
	fmt.Println("Part 2 : ", a2)
}

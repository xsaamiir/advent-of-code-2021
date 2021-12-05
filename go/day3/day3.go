package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Day3 struct{}

func (d Day3) InputGenerator(reader io.Reader) (interface{}, error) {
	scanner := bufio.NewScanner(reader)

	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return output, nil
}

func (d Day3) SolverPart1(v interface{}) (interface{}, error) {
	input := v.([]string)

	nInputs := len(input)
	nBits := len(input[0])

	matrix := make([][]int, nInputs)
	for i := 0; i < nInputs; i++ {
		matrix[i] = make([]int, nBits)
	}

	for i := 0; i < nInputs; i++ {
		for j := 0; j < nBits; j++ {
			matrix[i][j] = int(input[i][j] - '0')
		}
	}

	var gammaArray []uint

	for i := 0; i < nBits; i++ {
		var mostCommon uint
		var numberOfOnes int
		var numberOfZeros int

		for j := 0; j < nInputs; j++ {
			b := matrix[j][i]

			if b == 0 {
				numberOfZeros++
			} else {
				numberOfOnes++
			}
		}
		// end of vertical iter

		if numberOfOnes > numberOfZeros {
			mostCommon = 1
		} else {
			mostCommon = 0
		}
		gammaArray = append(gammaArray, mostCommon)
	}

	gammaRate := arrayToBit(gammaArray)
	epsilonRate := arrayToBit(invert(gammaArray))

	return gammaRate * epsilonRate, nil
}

func invert(b []uint) []uint {
	inverted := make([]uint, len(b))
	for i, n := range b {
		if n == 0 {
			inverted[i] = 1
		} else {
			inverted[i] = 0
		}
	}

	return inverted
}

func arrayToBit(is []uint) uint {
	var sb strings.Builder

	for _, i := range is {
		sb.WriteString(strconv.Itoa(int(i)))
	}

	n, _ := strconv.ParseInt(sb.String(), 2, 64)

	return uint(n)
}

func (d Day3) SolverPart2(v interface{}) (interface{}, error) {
	input := v.([]string)

	nInputs := len(input)
	nBits := len(input[0])

	matrix := make([][]uint, nInputs)
	for i := 0; i < nInputs; i++ {
		matrix[i] = make([]uint, nBits)
	}

	for i := 0; i < nInputs; i++ {
		for j := 0; j < nBits; j++ {
			matrix[i][j] = uint(input[i][j] - '0')
		}
	}

	var i uint
	left := matrix

	for {
		left = iter(left, i)
		i++

		if len(left) == 1 {
			break
		}
	}

	oxygenGeneratorRating := arrayToBit(left[0])

	var j uint
	left2 := matrix

	for {
		left2 = iter2(left2, j)
		j++

		if len(left2) == 1 {
			break
		}
	}

	co2ScrubberRating := arrayToBit(left2[0])

	return oxygenGeneratorRating * co2ScrubberRating, nil
}

func iter(m [][]uint, p uint) [][]uint {
	var (
		nZeros uint
		nOnes  uint
		zeros  [][]uint
		ones   [][]uint
	)

	for _, line := range m {
		b := line[p]
		if b == 0 {
			nZeros++
			zeros = append(zeros, line)
		} else {
			nOnes++
			ones = append(ones, line)
		}
	}

	if nZeros > nOnes {
		return zeros
	} else {
		return ones
	}
}

func iter2(m [][]uint, p uint) [][]uint {
	var (
		nZeros uint
		nOnes  uint
		zeros  [][]uint
		ones   [][]uint
	)

	for _, line := range m {
		b := line[p]
		if b == 0 {
			nZeros++
			zeros = append(zeros, line)
		} else {
			nOnes++
			ones = append(ones, line)
		}
	}

	if nZeros <= nOnes {
		return zeros
	} else {
		return ones
	}
}

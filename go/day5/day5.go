package day5

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Day5 struct{}

type Line struct {
	Start [2]int
	End   [2]int
}

func (l Line) X1() int {
	return l.Start[0]
}

func (l Line) Y1() int {
	return l.Start[1]
}

func (l Line) X2() int {
	return l.End[0]
}

func (l Line) Y2() int {
	return l.End[1]
}

func (l Line) IsHorizontal() (bool, int, int) {
	if l.Y1() != l.Y2() {
		return false, 0, 0
	}

	xStart := l.X1()
	xEnd := l.X2()

	if l.X1() > l.X2() {
		xStart, xEnd = xEnd, xStart
	}

	return true, xStart, xEnd
}

func (l Line) IsVertical() (bool, int, int) {
	if l.X1() != l.X2() {
		return false, 0, 0
	}

	yStart := l.Y1()
	yEnd := l.Y2()

	if l.Y1() > l.Y2() {
		yStart, yEnd = yEnd, yStart
	}

	return true, yStart, yEnd
}

func (l Line) IsDiagonal() (bool, [][2]int) {
	var cs [][2]int

	slope := float64(l.Y2()-l.Y1()) / float64(l.X2()-l.X1())

	is45 := math.Abs(slope) == 1
	if !is45 {
		return false, nil
	}

	diff := l.X1() - l.X2()
	sign := diff / int(math.Abs(float64(diff)))

	for i := 0; i <= int(math.Abs(float64(diff))); i++ {
		p := [2]int{l.X1() + (i * -sign), l.Y1() + ((i * -sign) * int(slope))}
		cs = append(cs, p)
	}

	return true, cs
}

func (d Day5) InputGenerator(reader io.Reader) (interface{}, error) {
	scanner := bufio.NewScanner(reader)

	var output []Line

	for scanner.Scan() {
		text := scanner.Text()

		cs := strings.Split(text, "->")
		if l := len(cs); l != 2 {
			return nil, fmt.Errorf("error splitting on '->', expected two, got %d", l)
		}

		start, end := cs[0], cs[1]

		c1 := strings.Split(start, ",")
		c2 := strings.Split(end, ",")

		if len(c1) != 2 && len(c2) != 2 {
			return nil, fmt.Errorf("error splitting on ',', expected two")
		}

		x1, err := strconv.Atoi(strings.TrimSpace(c1[0]))
		if err != nil {
			return nil, fmt.Errorf("error parsing x1 as string: %w", err)
		}

		y1, err := strconv.Atoi(strings.TrimSpace(c1[1]))
		if err != nil {
			return nil, fmt.Errorf("error parsing y1 as string: %w", err)
		}

		x2, err := strconv.Atoi(strings.TrimSpace(c2[0]))
		if err != nil {
			return nil, fmt.Errorf("error parsing x2 as string: %w", err)
		}

		y2, err := strconv.Atoi(strings.TrimSpace(c2[1]))
		if err != nil {
			return nil, fmt.Errorf("error parsing y2 as string: %w", err)
		}

		output = append(output, Line{
			Start: [2]int{x1, y1},
			End:   [2]int{x2, y2},
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return output, nil
}

type Grid struct {
	Lines []Line
}

func NewGrid(lines []Line) Grid {
	return Grid{Lines: lines}
}

func (g Grid) Covered(includeDiagonal bool) [][]int {
	lines := g.Lines
	xs, ys := g.dimensions()

	covered := make([][]int, ys)
	for i := range covered {
		covered[i] = make([]int, xs)
	}

	for _, line := range lines {
		if ok, s, e := line.IsHorizontal(); ok {
			for i := s; i <= e; i++ {
				covered[line.Y1()][i]++
			}
		}

		if ok, s, e := line.IsVertical(); ok {
			for i := s; i <= e; i++ {
				covered[i][line.X1()]++
			}
		}

		if ok, cs := line.IsDiagonal(); includeDiagonal && ok {
			for _, c := range cs {
				covered[c[1]][c[0]]++
			}
		}
	}

	return covered
}

func (g Grid) dimensions() (int, int) {
	var (
		x int
		y int
	)

	for _, line := range g.Lines {
		if line.X1() > x {
			x = line.X1()
		}

		if line.X2() > x {
			x = line.X2()
		}

		if line.Y1() > y {
			y = line.Y1()
		}

		if line.Y2() > y {
			y = line.Y2()
		}
	}

	return x + 1, y + 1
}

func (g Grid) String() string {
	var sb strings.Builder

	for i, lines := range g.Covered(false) {
		for _, point := range lines {
			if point == 0 {
				sb.WriteString(".")
			} else {
				sb.WriteString(strconv.Itoa(point))
			}
		}

		if i != len(lines)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func (g Grid) Overlaps(includeDiagonal bool) int {
	var count int

	for _, lines := range g.Covered(includeDiagonal) {
		for _, p := range lines {
			if p >= 2 {
				count++
			}
		}
	}

	return count
}

func (d Day5) SolverPart1(v interface{}) (interface{}, error) {
	input := v.([]Line)

	grid := NewGrid(input)

	return grid.Overlaps(false), nil
}

func (d Day5) SolverPart2(v interface{}) (interface{}, error) {
	input := v.([]Line)

	grid := NewGrid(input)

	return grid.Overlaps(true), nil
}

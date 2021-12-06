package runner

import (
	"fmt"
	"io"
	"os"
)

type Day interface {
	// InputGenerator takes the input from the input file and should return
	// the data type expected by the solver functions.
	InputGenerator(io.Reader) (interface{}, error)
	// SolverPart1 takes as an input the output of InputGenerator.
	SolverPart1(interface{}) (interface{}, error)
	// SolverPart2 takes as an input the output of InputGenerator.
	SolverPart2(interface{}) (interface{}, error)
}

type DayRunner struct {
	Number int
	Day
}

func (d *DayRunner) Run() error {
	if d.Number == 0 {
		return fmt.Errorf("day number is not set")
	}

	f, err := os.Open(fmt.Sprintf("./day%d/day%d.txt", d.Number, d.Number))
	if err != nil {
		return err
	}

	defer f.Close()

	i, err := d.InputGenerator(f)
	if err != nil {
		return err
	}

	a1, err := d.SolverPart1(i)
	if err != nil {
		return err
	}

	fmt.Println("Part 1 : ", a1)

	a2, err := d.SolverPart2(i)
	if err != nil {
		return err
	}

	fmt.Println("Part 2 : ", a2)

	return nil
}

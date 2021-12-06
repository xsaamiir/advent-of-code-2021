package main

import (
	"github.com/sharkyze/advent-of-code/go/runner"
)

func main() {
	r := runner.DayRunner{
		Number: 6,
		Day:    Day6{},
	}

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

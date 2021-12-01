package main

import (
	"fmt"
	"os"
)

func main() {
	day := Day1{}
	dr := DayRunner{Day: day, Number: 1}

	if err := dr.Run(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}

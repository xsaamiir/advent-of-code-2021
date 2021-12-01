package aoc

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_inputGeneratorDay1(t *testing.T) {
	const input = "5\n7\n49"
	got, err := intputGeneratorDay1(bytes.NewBufferString(input))
	if err != nil {
		t.Errorf(err.Error())
	}

	if !cmp.Equal([]int{5, 7, 49}, got) {
		t.Fatalf(cmp.Diff([]int{5, 7, 49}, got))
	}
}

func Test_solverDay1Part1(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	got := solverDay1Part1(input)

	if got != 7 {
		t.Errorf("Expected = 7, got = %d", got)
	}
}

func Test_solverDay1Part2(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	got := solverDay1Part2(input)

	if got != 5 {
		t.Errorf("Expected = 5, got = %d", got)
	}
}

func Test_mainDay1Part1(t *testing.T) {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	mainDay1Part1()
}

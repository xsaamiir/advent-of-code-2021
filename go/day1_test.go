package main

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDay1_InputGenerator(t *testing.T) {
	const input = "5\n7\n49"
	got, err := Day1{}.InputGenerator(bytes.NewBufferString(input))
	if err != nil {
		t.Errorf(err.Error())
	}

	if !cmp.Equal([]int{5, 7, 49}, got) {
		t.Fatalf(cmp.Diff([]int{5, 7, 49}, got))
	}
}

func TestDay1_SolverPart1(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	got, err := Day1{}.SolverPart1(input)
	if err != nil {
		t.Error()
	}

	if got != 7 {
		t.Errorf("Expected = 7, got = %d", got)
	}
}

func TestDay1_SolverPart2(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	got, err := Day1{}.SolverPart2(input)
	if err != nil {
		t.Error()
	}

	if got != 5 {
		t.Errorf("Expected = 5, got = %d", got)
	}
}

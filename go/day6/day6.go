package main

import (
	"io"
	"strconv"
	"strings"
)

type Day6 struct{}

type Lanternfish struct {
	Timer int
}

func (l *Lanternfish) Tick() (reproduced bool, baby Lanternfish) {
	l.Timer--
	if l.Timer >= 0 {
		return false, Lanternfish{}
	}

	l.Timer = 6

	return true, Lanternfish{Timer: 8}
}

type LanternfishSchool struct {
	Lanternfish []Lanternfish
}

func (s *LanternfishSchool) Tick(days int) {
	for i := 1; i <= days; i++ {
		for j, lanternfish := range s.Lanternfish {
			reproduced, nl := lanternfish.Tick()
			if reproduced {
				s.Lanternfish = append(s.Lanternfish, nl)
			}

			s.Lanternfish[j] = lanternfish
		}
	}
}

func (s *LanternfishSchool) SizeAfter(nDays int) int {
	var days [9]int

	for _, lanternfish := range s.Lanternfish {
		days[lanternfish.Timer]++
	}

	for i := 0; i < nDays; i++ {
		var next [9]int

		for i := 1; i < 9; i++ {
			next[i-1] = days[i]
		}

		next[6] += days[0]
		next[8] += days[0]

		days = next
	}

	var sum int
	for _, day := range days {
		sum += day
	}

	return sum
}

func (s LanternfishSchool) Size() int {
	return len(s.Lanternfish)
}

func (d Day6) InputGenerator(reader io.Reader) (interface{}, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	ss := strings.Split(strings.TrimSpace(string(b)), ",")

	var fish []Lanternfish

	for _, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		fish = append(fish, Lanternfish{Timer: n})
	}

	return LanternfishSchool{Lanternfish: fish}, nil
}

func (d Day6) SolverPart1(v interface{}) (interface{}, error) {
	input := v.(LanternfishSchool)

	return input.SizeAfter(80), nil
}

func (d Day6) SolverPart2(v interface{}) (interface{}, error) {
	input := v.(LanternfishSchool)

	return input.SizeAfter(256), nil
}

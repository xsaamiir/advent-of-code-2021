package day4

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Day4 struct{}

type Board struct {
	Grid            [5][5]int
	MarkedPositions [5][5]bool
	LastMarked      int
}

func NewBoardFromString(s string) (Board, error) {
	var b [5][5]int

	lines := strings.Split(strings.TrimSuffix(s, "\n"), "\n")
	if len(lines) != 5 {
		return Board{}, fmt.Errorf("error number of lines is not 5: %#v", lines)
	}

	for i, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) != 5 {
			return Board{}, fmt.Errorf("error number of number per line is not 5: %#v", numbers)
		}

		for j, number := range numbers {
			n, err := strconv.Atoi(number)
			if err != nil {
				return Board{}, err
			}

			b[i][j] = n
		}
	}

	return Board{
		Grid:            b,
		MarkedPositions: [5][5]bool{},
	}, nil
}

func (b Board) Mark(n int) Board {
	marked := b

	for i, line := range b.Grid {
		for j, pos := range line {
			if pos == n {
				marked.MarkedPositions[i][j] = true
				marked.LastMarked = n
			}
		}
	}

	return marked
}

// HasWon returns true if a board has at least one complete row or column of marked numbers.
func (b Board) HasWon() (bool, int) {
	var hasWon bool

	for _, line := range b.MarkedPositions {
		if all(line) {
			hasWon = true
			break
		}
	}

	t := transpose(b.MarkedPositions)

	for _, column := range t {
		if all(column) {
			hasWon = true
			break
		}
	}

	if !hasWon {
		return false, 0
	}

	return true, b.score()
}

func (b Board) score() int {
	var sum int

	for i, line := range b.MarkedPositions {
		for j, pos := range line {
			if !pos {
				sum += b.Grid[i][j]
			}
		}
	}

	return sum * b.LastMarked
}

func transpose(grid [5][5]bool) [5][5]bool {
	var t [5][5]bool

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			t[j][i] = grid[i][j]
		}
	}

	return t
}

func all(ps [5]bool) bool {
	for _, p := range ps {
		if !p {
			return false
		}
	}

	return true
}

type Game struct {
	Boards       []Board
	DrawnNumbers []int
}

type GameResult struct {
	IsGameOver bool
	BoardIndex int
	Score      int
}

// Play iterates over DrawnNumbers and returns when the last board wins.
// It returns whether the game is over, the index of the winning board
// and the winning score.
func (g Game) Play() []GameResult {
	var winners []GameResult

	for _, number := range g.DrawnNumbers {
		for j, board := range g.Boards {
			if hasWon, _ := g.Boards[j].HasWon(); hasWon {
				continue
			}

			var r GameResult

			g.Boards[j] = board.Mark(number)
			r.IsGameOver, r.Score = g.Boards[j].HasWon()
			if r.IsGameOver {
				r.BoardIndex = j
				winners = append(winners, r)
			}
		}
	}

	return winners
}

func parseDrawnNumbers(s string) ([]int, error) {
	var drawnNumbers []int

	ss := strings.Split(s, ",")

	for _, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		drawnNumbers = append(drawnNumbers, n)
	}

	return drawnNumbers, nil
}

func (d Day4) InputGenerator(reader io.Reader) (interface{}, error) {
	scanner := bufio.NewScanner(reader)

	var (
		game Game
		err  error
	)

	// The first line are the list of drawn numbers.
	scanner.Scan()
	game.DrawnNumbers, err = parseDrawnNumbers(scanner.Text())
	if err != nil {
		return nil, err
	}

	var (
		sb strings.Builder
		i  int
	)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.TrimSpace(text) == "" {
			i = 0
			sb.Reset()
			continue
		}

		sb.WriteString(text)
		sb.WriteString("\n")
		i++

		if i == 5 {
			board, err := NewBoardFromString(sb.String())
			if err != nil {
				return nil, err
			}

			game.Boards = append(game.Boards, board)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return game, nil
}

func (d Day4) SolverPart1(v interface{}) (interface{}, error) {
	input := v.(Game)

	winners := input.Play()

	return winners[0], nil
}

func (d Day4) SolverPart2(v interface{}) (interface{}, error) {
	input := v.(Game)

	winners := input.Play()

	return winners[len(winners)-1], nil
}

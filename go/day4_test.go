package main

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBoardFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := map[string]struct {
		args    args
		want    Board
		wantErr bool
	}{
		"ok": {
			args: args{
				s: "22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n",
			},
			want: Board{
				Grid: [5][5]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewBoardFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBoardFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoardFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Mark(t *testing.T) {
	type fields struct {
		Grid            [5][5]int
		MarkedPositions [5][5]bool
	}

	type args struct {
		n int
	}

	tests := map[string]struct {
		fields fields
		args   args
		want   Board
	}{
		"number found": {
			fields: fields{
				Grid: [5][5]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				MarkedPositions: [5][5]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
			args: args{n: 14},
			want: Board{
				Grid: [5][5]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				MarkedPositions: [5][5]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, true, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
				LastMarked: 14,
			},
		},
		"number not found": {
			fields: fields{
				Grid: [5][5]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				MarkedPositions: [5][5]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
			args: args{n: 41},
			want: Board{
				Grid: [5][5]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				MarkedPositions: [5][5]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b := Board{
				Grid:            tt.fields.Grid,
				MarkedPositions: tt.fields.MarkedPositions,
			}

			if got := b.Mark(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mark() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_HasWon(t *testing.T) {
	type fields struct {
		Grid            [5][5]int
		MarkedPositions [5][5]bool
		LastMarked      int
	}

	type want struct {
		won   bool
		score int
	}

	tests := map[string]struct {
		fields fields
		want   want
	}{
		"zero values": {
			fields: fields{
				Grid: [5][5]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				MarkedPositions: [5][5]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
			want: want{
				won:   false,
				score: 0,
			},
		},
		"nothing marked": {
			fields: fields{
				Grid: [5][5]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				MarkedPositions: [5][5]bool{
					{true, false, false, true, true},
					{false, false, true, false, false},
					{false, true, false, true, false},
					{false, false, false, false, true},
					{true, false, false, true, false},
				},
			},
			want: want{
				won:   false,
				score: 0,
			},
		},
		"lines wins": {
			fields: fields{
				Grid: [5][5]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				MarkedPositions: [5][5]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{true, true, true, true, true},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
			want: want{
				won:   true,
				score: 0,
			},
		},
		"column wins": {
			fields: fields{
				Grid: [5][5]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				MarkedPositions: [5][5]bool{
					{false, true, false, false, false},
					{false, true, false, false, false},
					{false, true, false, false, false},
					{false, true, false, false, false},
					{false, true, false, false, false},
				},
			},
			want: want{
				won:   true,
				score: 0,
			},
		},
		"sample input wins": {
			fields: fields{
				Grid: [5][5]int{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
				MarkedPositions: [5][5]bool{
					{true, true, true, true, true},
					{false, false, false, true, false},
					{false, false, true, false, false},
					{false, true, false, false, true},
					{true, true, false, false, true},
				},
				LastMarked: 24,
			},
			want: want{
				won:   true,
				score: 4512,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b := Board{
				Grid:            tt.fields.Grid,
				MarkedPositions: tt.fields.MarkedPositions,
				LastMarked:      tt.fields.LastMarked,
			}

			got, score := b.HasWon()

			if !cmp.Equal(tt.want.won, got) {
				t.Error(cmp.Diff(tt.want.won, got))
			}

			if !cmp.Equal(tt.want.score, score) {
				t.Error(cmp.Diff(tt.want.score, score))
			}
		})
	}
}

func Test_parseDrawnNumbers(t *testing.T) {
	type args struct {
		s string
	}
	tests := map[string]struct {
		args    args
		want    []int
		wantErr bool
	}{
		"ok": {
			args: args{
				s: "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
			},
			want:    []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parseDrawnNumbers(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDrawnNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDrawnNumbers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_Play(t *testing.T) {
	type fields struct {
		Boards       []Board
		DrawnNumbers []int
	}

	type want struct {
		GameResult []GameResult
	}

	tests := map[string]struct {
		fields fields
		want   want
	}{
		"ok": {
			fields: fields{
				Boards: []Board{
					{
						Grid: [5][5]int{
							{22, 13, 17, 11, 0},
							{8, 2, 23, 4, 24},
							{21, 9, 14, 16, 7},
							{6, 10, 3, 18, 5},
							{1, 12, 20, 15, 19},
						},
						MarkedPositions: [5][5]bool{},
					},
					{
						Grid: [5][5]int{
							{3, 15, 0, 2, 22},
							{9, 18, 13, 17, 5},
							{19, 8, 7, 25, 23},
							{20, 11, 10, 24, 4},
							{14, 21, 16, 12, 6},
						},
					},
					{
						Grid: [5][5]int{
							{14, 21, 17, 24, 4},
							{10, 16, 15, 9, 19},
							{18, 8, 23, 26, 20},
							{22, 11, 13, 6, 5},
							{2, 0, 12, 3, 7},
						},
					},
				},
				DrawnNumbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			},
			want: want{
				GameResult: []GameResult{
					{
						IsGameOver: true,
						BoardIndex: 2,
						Score:      4512,
					},
					{
						IsGameOver: true,
						BoardIndex: 0,
						Score:      2192,
					},
					{
						IsGameOver: true,
						BoardIndex: 1,
						Score:      1924,
					},
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := Game{
				Boards:       tt.fields.Boards,
				DrawnNumbers: tt.fields.DrawnNumbers,
			}

			result := g.Play()

			if !cmp.Equal(tt.want.GameResult, result) {
				t.Error(cmp.Diff(tt.want.GameResult, result))
			}
		})
	}
}

func TestDay4_InputGenerator(t *testing.T) {
	type args struct {
		input string
	}

	tests := map[string]struct {
		args args
		want interface {
		}
		wantErr bool
	}{
		"ok": {
			args: args{
				input: "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1\n\n22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n\n 3 15  0  2 22\n 9 18 13 17  5\n19  8  7 25 23\n20 11 10 24  4\n14 21 16 12  6\n\n14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7",
			},
			want: Game{
				Boards: []Board{
					{
						Grid: [5][5]int{
							{22, 13, 17, 11, 0},
							{8, 2, 23, 4, 24},
							{21, 9, 14, 16, 7},
							{6, 10, 3, 18, 5},
							{1, 12, 20, 15, 19},
						},
						MarkedPositions: [5][5]bool{},
					},
					{
						Grid: [5][5]int{
							{3, 15, 0, 2, 22},
							{9, 18, 13, 17, 5},
							{19, 8, 7, 25, 23},
							{20, 11, 10, 24, 4},
							{14, 21, 16, 12, 6},
						},
					},
					{
						Grid: [5][5]int{
							{14, 21, 17, 24, 4},
							{10, 16, 15, 9, 19},
							{18, 8, 23, 26, 20},
							{22, 11, 13, 6, 5},
							{2, 0, 12, 3, 7},
						},
					},
				},
				DrawnNumbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day4{}
			got, err := d.InputGenerator(bytes.NewBufferString(tt.args.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("InputGenerator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(tt.want, got) {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestDay4_SolverPart1(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := map[string]struct {
		args    args
		want    interface{}
		wantErr bool
	}{
		"ok": {
			args: args{
				v: Game{
					Boards: []Board{
						{
							Grid: [5][5]int{
								{22, 13, 17, 11, 0},
								{8, 2, 23, 4, 24},
								{21, 9, 14, 16, 7},
								{6, 10, 3, 18, 5},
								{1, 12, 20, 15, 19},
							},
							MarkedPositions: [5][5]bool{},
						},
						{
							Grid: [5][5]int{
								{3, 15, 0, 2, 22},
								{9, 18, 13, 17, 5},
								{19, 8, 7, 25, 23},
								{20, 11, 10, 24, 4},
								{14, 21, 16, 12, 6},
							},
						},
						{
							Grid: [5][5]int{
								{14, 21, 17, 24, 4},
								{10, 16, 15, 9, 19},
								{18, 8, 23, 26, 20},
								{22, 11, 13, 6, 5},
								{2, 0, 12, 3, 7},
							},
						},
					},
					DrawnNumbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				},
			},
			want: GameResult{
				IsGameOver: true,
				BoardIndex: 2,
				Score:      4512,
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day4{}
			got, err := d.SolverPart1(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolverPart1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(tt.want, got) {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestDay4_SolverPart2(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := map[string]struct {
		args    args
		want    interface{}
		wantErr bool
	}{
		"ok": {
			args: args{
				v: Game{
					Boards: []Board{
						{
							Grid: [5][5]int{
								{22, 13, 17, 11, 0},
								{8, 2, 23, 4, 24},
								{21, 9, 14, 16, 7},
								{6, 10, 3, 18, 5},
								{1, 12, 20, 15, 19},
							},
							MarkedPositions: [5][5]bool{},
						},
						{
							Grid: [5][5]int{
								{3, 15, 0, 2, 22},
								{9, 18, 13, 17, 5},
								{19, 8, 7, 25, 23},
								{20, 11, 10, 24, 4},
								{14, 21, 16, 12, 6},
							},
						},
						{
							Grid: [5][5]int{
								{14, 21, 17, 24, 4},
								{10, 16, 15, 9, 19},
								{18, 8, 23, 26, 20},
								{22, 11, 13, 6, 5},
								{2, 0, 12, 3, 7},
							},
						},
					},
					DrawnNumbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				},
			},
			want: GameResult{
				IsGameOver: true,
				BoardIndex: 1,
				Score:      1924,
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day4{}
			got, err := d.SolverPart2(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolverPart2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(tt.want, got) {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}

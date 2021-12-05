package main

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLine_IsDiagonal(t *testing.T) {
	type fields struct {
		Start [2]int
		End   [2]int
	}

	tests := map[string]struct {
		fields fields
		want   bool
		want1  [][2]int
	}{
		"ok 1": {
			fields: fields{
				Start: [2]int{1, 1},
				End:   [2]int{3, 3},
			},
			want:  true,
			want1: [][2]int{{1, 1}, {2, 2}, {3, 3}},
		},
		"ok 2": {
			fields: fields{
				Start: [2]int{3, 3},
				End:   [2]int{1, 1},
			},
			want:  true,
			want1: [][2]int{{3, 3}, {2, 2}, {1, 1}},
		},
		"ok 3": {
			fields: fields{
				Start: [2]int{9, 7},
				End:   [2]int{7, 9},
			},
			want:  true,
			want1: [][2]int{{9, 7}, {8, 8}, {7, 9}},
		},
		"ok 4": {
			fields: fields{
				Start: [2]int{7, 9},
				End:   [2]int{9, 7},
			},
			want:  true,
			want1: [][2]int{{7, 9}, {8, 8}, {9, 7}},
		},
		"ok 5": {
			fields: fields{
				Start: [2]int{5, 5},
				End:   [2]int{8, 2},
			},
			want:  true,
			want1: [][2]int{{5, 5}, {6, 4}, {7, 3}, {8, 2}},
		},
		"nok": {
			fields: fields{
				Start: [2]int{0, 9},
				End:   [2]int{5, 9},
			},
			want:  false,
			want1: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := Line{
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}

			got, got1 := l.IsDiagonal()
			if !cmp.Equal(tt.want, got) {
				t.Error(cmp.Diff(tt.want, got))
			}

			if !cmp.Equal(tt.want1, got1) {
				t.Error(cmp.Diff(tt.want1, got1))
			}
		})
	}
}

func TestDay5_InputGenerator(t *testing.T) {
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
				input: "0,9 -> 5,9\n8,0 -> 0,8\n9,4 -> 3,4\n2,2 -> 2,1\n7,0 -> 7,4\n6,4 -> 2,0\n0,9 -> 2,9\n3,4 -> 1,4\n0,0 -> 8,8\n5,5 -> 8,2\n",
			},
			want: []Line{
				{Start: [2]int{0, 9}, End: [2]int{5, 9}},
				{Start: [2]int{8, 0}, End: [2]int{0, 8}},
				{Start: [2]int{9, 4}, End: [2]int{3, 4}},
				{Start: [2]int{2, 2}, End: [2]int{2, 1}},
				{Start: [2]int{7, 0}, End: [2]int{7, 4}},
				{Start: [2]int{6, 4}, End: [2]int{2, 0}},
				{Start: [2]int{0, 9}, End: [2]int{2, 9}},
				{Start: [2]int{3, 4}, End: [2]int{1, 4}},
				{Start: [2]int{0, 0}, End: [2]int{8, 8}},
				{Start: [2]int{5, 5}, End: [2]int{8, 2}},
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day5{}
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

func TestGrid_Covered(t *testing.T) {
	type args struct {
		lines           []Line
		includeDiagonal bool
	}

	tests := map[string]struct {
		args args
		want [][]int
	}{
		"horizontal & vertical": {
			args: args{
				lines: []Line{
					{Start: [2]int{0, 9}, End: [2]int{5, 9}},
					{Start: [2]int{8, 0}, End: [2]int{0, 8}},
					{Start: [2]int{9, 4}, End: [2]int{3, 4}},
					{Start: [2]int{2, 2}, End: [2]int{2, 1}},
					{Start: [2]int{7, 0}, End: [2]int{7, 4}},
					{Start: [2]int{6, 4}, End: [2]int{2, 0}},
					{Start: [2]int{0, 9}, End: [2]int{2, 9}},
					{Start: [2]int{3, 4}, End: [2]int{1, 4}},
					{Start: [2]int{0, 0}, End: [2]int{8, 8}},
					{Start: [2]int{5, 5}, End: [2]int{8, 2}},
				},
				includeDiagonal: false,
			},
			want: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
			},
		},
		"diagonal": {
			args: args{
				lines: []Line{
					{Start: [2]int{0, 9}, End: [2]int{5, 9}},
					{Start: [2]int{8, 0}, End: [2]int{0, 8}},
					{Start: [2]int{9, 4}, End: [2]int{3, 4}},
					{Start: [2]int{2, 2}, End: [2]int{2, 1}},
					{Start: [2]int{7, 0}, End: [2]int{7, 4}},
					{Start: [2]int{6, 4}, End: [2]int{2, 0}},
					{Start: [2]int{0, 9}, End: [2]int{2, 9}},
					{Start: [2]int{3, 4}, End: [2]int{1, 4}},
					{Start: [2]int{0, 0}, End: [2]int{8, 8}},
					{Start: [2]int{5, 5}, End: [2]int{8, 2}},
				},
				includeDiagonal: true,
			},
			want: [][]int{
				{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
				{0, 1, 1, 1, 0, 0, 0, 2, 0, 0},
				{0, 0, 2, 0, 1, 0, 1, 1, 1, 0},
				{0, 0, 0, 1, 0, 2, 0, 2, 0, 0},
				{0, 1, 1, 2, 3, 1, 3, 2, 1, 1},
				{0, 0, 0, 1, 0, 2, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			grid := NewGrid(tt.args.lines)

			got := grid.Covered(tt.args.includeDiagonal)
			if !cmp.Equal(tt.want, got) {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestGrid_String(t *testing.T) {
	type fields struct {
		Lines []Line
	}

	tests := map[string]struct {
		fields fields
		want   string
	}{
		"ok": {
			fields: fields{
				Lines: []Line{
					{Start: [2]int{0, 9}, End: [2]int{5, 9}},
					{Start: [2]int{8, 0}, End: [2]int{0, 8}},
					{Start: [2]int{9, 4}, End: [2]int{3, 4}},
					{Start: [2]int{2, 2}, End: [2]int{2, 1}},
					{Start: [2]int{7, 0}, End: [2]int{7, 4}},
					{Start: [2]int{6, 4}, End: [2]int{2, 0}},
					{Start: [2]int{0, 9}, End: [2]int{2, 9}},
					{Start: [2]int{3, 4}, End: [2]int{1, 4}},
					{Start: [2]int{0, 0}, End: [2]int{8, 8}},
					{Start: [2]int{5, 5}, End: [2]int{8, 2}},
				},
			},
			want: `.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g := Grid{Lines: tt.fields.Lines}

			got := g.String()
			if !cmp.Equal(tt.want, got) {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestDay5_SolverPart1(t *testing.T) {
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
				v: []Line{
					{Start: [2]int{0, 9}, End: [2]int{5, 9}},
					{Start: [2]int{8, 0}, End: [2]int{0, 8}},
					{Start: [2]int{9, 4}, End: [2]int{3, 4}},
					{Start: [2]int{2, 2}, End: [2]int{2, 1}},
					{Start: [2]int{7, 0}, End: [2]int{7, 4}},
					{Start: [2]int{6, 4}, End: [2]int{2, 0}},
					{Start: [2]int{0, 9}, End: [2]int{2, 9}},
					{Start: [2]int{3, 4}, End: [2]int{1, 4}},
					{Start: [2]int{0, 0}, End: [2]int{8, 8}},
					{Start: [2]int{5, 5}, End: [2]int{8, 2}},
				},
			},
			want:    5,
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day5{}
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

func TestDay5_SolverPart2(t *testing.T) {
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
				v: []Line{
					{Start: [2]int{0, 9}, End: [2]int{5, 9}},
					{Start: [2]int{8, 0}, End: [2]int{0, 8}},
					{Start: [2]int{9, 4}, End: [2]int{3, 4}},
					{Start: [2]int{2, 2}, End: [2]int{2, 1}},
					{Start: [2]int{7, 0}, End: [2]int{7, 4}},
					{Start: [2]int{6, 4}, End: [2]int{2, 0}},
					{Start: [2]int{0, 9}, End: [2]int{2, 9}},
					{Start: [2]int{3, 4}, End: [2]int{1, 4}},
					{Start: [2]int{0, 0}, End: [2]int{8, 8}},
					{Start: [2]int{5, 5}, End: [2]int{8, 2}},
				},
			},
			want:    12,
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day5{}
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

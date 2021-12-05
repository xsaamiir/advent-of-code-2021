package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDay3_InputGenerator(t *testing.T) {
	type args struct {
		input string
	}

	tests := map[string]struct {
		args    args
		want    interface{}
		wantErr bool
	}{
		"ok": {
			args: args{
				input: "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010",
			},
			want:    []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day3{}
			got, err := d.InputGenerator(bytes.NewBufferString(tt.args.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("InputGenerator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InputGenerator() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay3_SolverPart1(t *testing.T) {
	type args struct {
		i interface{}
	}

	tests := map[string]struct {
		args    args
		want    interface{}
		wantErr bool
	}{
		"ok": {
			args: args{
				i: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			},
			want:    uint(198),
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day3{}
			got, err := d.SolverPart1(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolverPart1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolverPart1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay3_SolverPart2(t *testing.T) {
	type args struct {
		i interface{}
	}

	tests := map[string]struct {
		args    args
		want    interface{}
		wantErr bool
	}{
		"ok": {
			args: args{
				i: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			},
			want:    uint(230),
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day3{}
			got, err := d.SolverPart2(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolverPart2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolverPart2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iter(t *testing.T) {
	type args struct {
		m [][]uint
		p uint
	}

	tests := map[string]struct {
		args args
		want [][]uint
	}{
		"ok": {
			args: args{
				m: [][]uint{{0, 0, 1, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 1, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 1, 0, 1}, {0, 1, 1, 1, 1}, {0, 0, 1, 1, 1}, {1, 1, 1, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {0, 1, 0, 1, 0}},
				p: 0,
			},
			want: [][]uint{{1, 1, 1, 1, 0}, {1, 0, 1, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 1, 0, 1}, {1, 1, 1, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 1}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := iter(tt.args.m, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iter2(t *testing.T) {
	type args struct {
		m [][]uint
		p uint
	}

	tests := map[string]struct {
		args args
		want [][]uint
	}{
		"ok": {
			args: args{
				m: [][]uint{{0, 0, 1, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 1, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 1, 0, 1}, {0, 1, 1, 1, 1}, {0, 0, 1, 1, 1}, {1, 1, 1, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {0, 1, 0, 1, 0}},
				p: 0,
			},
			want: [][]uint{{0, 0, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 1, 1, 1}, {0, 0, 0, 1, 0}, {0, 1, 0, 1, 0}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := iter2(tt.args.m, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iter() = %v, want %v", got, tt.want)
			}
		})
	}
}

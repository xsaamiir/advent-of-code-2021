package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDay2_InputGenerator(t *testing.T) {
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
				input: "forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2",
			},
			want: []Step{
				{Direction: DirectionForward, Unit: 5},
				{Direction: DirectionDown, Unit: 5},
				{Direction: DirectionForward, Unit: 8},
				{Direction: DirectionUp, Unit: 3},
				{Direction: DirectionDown, Unit: 8},
				{Direction: DirectionForward, Unit: 2},
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day2{}
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

func TestDay2_SolverPart1(t *testing.T) {
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
				i: []Step{
					{Direction: DirectionForward, Unit: 5},
					{Direction: DirectionDown, Unit: 5},
					{Direction: DirectionForward, Unit: 8},
					{Direction: DirectionUp, Unit: 3},
					{Direction: DirectionDown, Unit: 8},
					{Direction: DirectionForward, Unit: 2},
				},
			},
			want:    150,
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day2{}
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

func TestDay2_SolverPart2(t *testing.T) {
	type args struct {
		i interface{}
	}

	tests := map[string]struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		"ok": {
			args: args{
				i: []Step{
					{Direction: DirectionForward, Unit: 5},
					{Direction: DirectionDown, Unit: 5},
					{Direction: DirectionForward, Unit: 8},
					{Direction: DirectionUp, Unit: 3},
					{Direction: DirectionDown, Unit: 8},
					{Direction: DirectionForward, Unit: 2},
				},
			},
			want:    900,
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day2{}
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

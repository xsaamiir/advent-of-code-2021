package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestDay6_InputGenerator(t *testing.T) {
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
				input: "3,4,3,1,2",
			},
			want: LanternfishSchool{
				Lanternfish: []Lanternfish{
					{Timer: 3},
					{Timer: 4},
					{Timer: 3},
					{Timer: 1},
					{Timer: 2},
				},
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day6{}

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

func TestDay6_SolverPart1(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := map[string]struct {
		args    args
		want    interface{}
		wantErr bool
	}{
		"80 days": {
			args: args{
				v: LanternfishSchool{
					Lanternfish: []Lanternfish{
						{Timer: 3},
						{Timer: 4},
						{Timer: 3},
						{Timer: 1},
						{Timer: 2},
					},
				},
			},
			want:    5934,
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day6{}

			got2 := numberOfFishAfter(tt.args.v.(LanternfishSchool).Lanternfish, 80)
			fmt.Println(got2)

			got, err := d.SolverPart1(tt.args.v)
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

func TestDay6_SolverPart2(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := map[string]struct {
		args    args
		want    interface{}
		wantErr bool
	}{
		"256 days": {
			args: args{
				v: LanternfishSchool{
					Lanternfish: []Lanternfish{
						{Timer: 3},
						{Timer: 4},
						{Timer: 3},
						{Timer: 1},
						{Timer: 2},
					},
				},
			},
			want:    26984457539,
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day6{}

			got, err := d.SolverPart2(tt.args.v)
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

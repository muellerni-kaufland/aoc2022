package day01

import (
	_ "embed"
	"io"
	"strings"
	"testing"
)

//go:embed input.txt
var actualInput string

var exampleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestTaskA(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example input", args{strings.NewReader(exampleInput)}, 24000},
		{"actual input", args{strings.NewReader(actualInput)}, 75622},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TaskA(tt.args.r); got != tt.want {
				t.Errorf("TaskA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskB(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example input", args{strings.NewReader(exampleInput)}, 45000},
		{"actual input", args{strings.NewReader(actualInput)}, 213159},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TaskB(tt.args.r); got != tt.want {
				t.Errorf("TaskB() = %v, want %v", got, tt.want)
			}
		})
	}
}

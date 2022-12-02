package day02

import (
	_ "embed"
	"io"
	"strings"
	"testing"
)

//go:embed input.txt
var actualInput string

const exampleInput = `A Y
B X
C Z`

func TestTaskA(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example input", args{strings.NewReader(exampleInput)}, 15},
		{"actual input", args{strings.NewReader(actualInput)}, 11063},
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
		{"example input", args{strings.NewReader(exampleInput)}, 12},
		{"actual input", args{strings.NewReader(actualInput)}, 10349},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TaskB(tt.args.r); got != tt.want {
				t.Errorf("TaskB() = %v, want %v", got, tt.want)
			}
		})
	}
}

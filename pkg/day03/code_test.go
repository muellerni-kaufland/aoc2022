package day03

import (
	_ "embed"
	"io"
	"strings"
	"testing"
)

//go:embed input.txt
var actualInput string

var exampleInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestTaskA(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example input", args{strings.NewReader(exampleInput)}, 157},
		{"actual input", args{strings.NewReader(actualInput)}, 8349},
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
		{"example input", args{strings.NewReader(exampleInput)}, 70},
		{"actual input", args{strings.NewReader(actualInput)}, 2681},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TaskB(tt.args.r); got != tt.want {
				t.Errorf("TaskB() = %v, want %v", got, tt.want)
			}
		})
	}
}

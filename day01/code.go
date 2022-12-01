package day01

import (
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`((\d+\n?)+)`)

// TaskA finds the maximum number of calories per input. If there are errors, it will return -1.
func TaskA(r io.Reader) int {
	matches := setup(r)
	if matches == nil {
		return -1
	}
	return computeCalories(matches, 1)
}

// TaskB finds the sum of the three maximum number of calories per input. If there are errors, it will return -1.
func TaskB(r io.Reader) int {
	matches := setup(r)
	if matches == nil {
		return -1
	}
	return computeCalories(matches, 3)
}

// computeCalories computes the sum of the `index` highest total calories carried by elves.
func computeCalories(matches []string, index int) int {
	var totalCalories []int
	for _, match := range matches {
		var calories int
		// strings.TrimSuffix is necessary to remove the trailing newline
		entries := strings.Split(strings.TrimSuffix(match, "\n"), "\n")
		for _, entry := range entries {
			cal, err := strconv.Atoi(entry)
			if err != nil {
				return -1
			}
			calories += cal
		}
		totalCalories = append(totalCalories, calories)
	}
	sort.Ints(totalCalories)
	var sum int
	for idx := index; idx > 0; idx-- {
		sum += totalCalories[len(totalCalories)-idx]
	}
	return sum
}

// setup loads the contents from `r` and matches them against `re`.
func setup(r io.Reader) []string {
	var builder strings.Builder
	_, err := io.Copy(&builder, r)
	if err != nil {
		return nil
	}
	matches := re.FindAllString(builder.String(), -1)
	if matches == nil {
		return nil
	}
	return matches
}

package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

func TestWindowIsUnique(t *testing.T) {
	windowSize := 4

	tcs := []struct {
		stream inputType
		want   bool
	}{
		{"fo", false},
		{"fooo", false},
		{"fghj", true},
		{"fghjf", true},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, windowIsUnique(tc.stream, windowSize))
		})
	}
}

func TestFindMarket(t *testing.T) {
	windowSize := 4

	tcs := []struct {
		stream inputType
		want   int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, findMarker(tc.stream, windowSize))
		})
	}
}

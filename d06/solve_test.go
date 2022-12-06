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
	type testcase struct {
		stream inputType
		want   int
	}

	t.Run("Test when window size is 4", func(t *testing.T) {
		windowSize := 4

		tcs := []testcase{
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
	})

	t.Run("Test when window size is 14", func(t *testing.T) {
		windowSize := 14

		tcs := []testcase{
			{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
			{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
			{"nppdvjthqldpwncqszvftbrmjlhg", 23},
			{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
			{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
		}

		for _, tc := range tcs {
			t.Run("", func(t *testing.T) {
				utils.Assert(t, tc.want, findMarker(tc.stream, windowSize))
			})
		}
	})

}

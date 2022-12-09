package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example = parseFile("example.txt")

func TestIsTouching(t *testing.T) {
	tcs := []struct {
		head position
		tail position
		want bool
	}{
		{position{0, 0}, position{0, 0}, true},
		{position{0, 0}, position{0, 1}, true},
		{position{0, 0}, position{1, 0}, true},
		{position{0, 0}, position{-1, 0}, true},
		{position{0, 0}, position{1, 1}, true},
		{position{0, 0}, position{1, 2}, false},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, isTouching(tc.head, tc.tail))
		})
	}
}

func TestCalcTailMove(t *testing.T) {
	tcs := []struct {
		head position
		tail position
		want position
	}{
		{position{0, 2}, position{0, 0}, position{0, 1}},
		{position{1, 2}, position{0, 0}, position{1, 1}},
		{position{0, 0}, position{2, 1}, position{-1, -1}},
		{position{3, 2}, position{1, 1}, position{1, 1}},
		{position{2, 3}, position{1, 1}, position{1, 1}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, calcTailMove(tc.head, tc.tail))
		})
	}
}

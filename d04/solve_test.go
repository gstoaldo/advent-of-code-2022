package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example = parseFile("example.txt")

func TestIsFullyContained(t *testing.T) {
	tcs := []struct {
		arg  []int
		want bool
	}{
		{example[0], false},
		{example[1], false},
		{example[2], false},
		{example[3], true},
		{example[4], true},
		{example[5], false},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := isFullyContained(tc.arg)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestIsOverlap(t *testing.T) {
	tcs := []struct {
		arg  []int
		want bool
	}{
		{example[0], false},
		{example[1], false},
		{example[2], true},
		{example[3], true},
		{example[4], true},
		{example[5], true},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := isOverlap(tc.arg)
			utils.Assert(t, tc.want, got)
		})
	}
}

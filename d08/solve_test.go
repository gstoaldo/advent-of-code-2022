package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example = parseFile("example.txt")

func TestIsOnTheEdge(t *testing.T) {
	tcs := []struct {
		grid inputType
		i    int
		j    int
		want bool
	}{
		{example, 0, 0, true},
		{example, 0, 1, true},
		{example, 1, 1, false},
		{example, 4, 4, true},
		{example, 3, 3, false},
		{example, 1, 4, true},
	}

	for _, tc := range tcs {
		utils.Assert(t, tc.want, isOnTheEdge(tc.grid, tc.i, tc.j))
	}
}

func TestIsVisible(t *testing.T) {
	tcs := []struct {
		grid inputType
		i    int
		j    int
		want bool
	}{
		{example, 1, 1, true},
		{example, 1, 2, true},
		{example, 1, 3, false},
		{example, 2, 1, true},
		{example, 2, 2, false},
		{example, 2, 3, true},
		{example, 3, 1, false},
		{example, 3, 2, true},
		{example, 3, 3, false},
	}

	for _, tc := range tcs {
		utils.Assert(t, tc.want, isVisible(tc.grid, tc.i, tc.j))
	}
}

func TestGetScenicScore(t *testing.T) {
	tcs := []struct {
		grid inputType
		i    int
		j    int
		want int
	}{
		{example, 1, 2, 4},
		{example, 3, 2, 8},
	}

	for _, tc := range tcs {
		utils.Assert(t, tc.want, getScenicScore(tc.grid, tc.i, tc.j))
	}
}

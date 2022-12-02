package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

func TestGetRoundScore(t *testing.T) {
	tcs := []struct {
		opShape string
		myShape string
		want    int
	}{
		{"A", "Y", 6},
		{"B", "X", 0},
		{"C", "Z", 3},
		{"A", "Z", 0},
		{"C", "X", 6},
	}

	for _, tc := range tcs {
		got := getRoundScore(tc.opShape, mapColumn[tc.myShape])
		utils.Assert(t, tc.want, got)
	}
}

func TestGetTotalScore(t *testing.T) {
	rounds := parseFile("example.txt")
	got := getTotalScore(rounds)
	utils.Assert(t, 15, got)
}

func TestGetShapeToResult(t *testing.T) {
	tcs := []struct {
		opShape string
		result  string
		want    string
	}{
		{"A", "Y", "A"},
		{"B", "X", "A"},
		{"C", "Z", "A"},
		{"A", "X", "C"},
	}

	for _, tc := range tcs {
		got := getShapeToResult(tc.opShape, tc.result)
		utils.Assert(t, tc.want, got)
	}
}

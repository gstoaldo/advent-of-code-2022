package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example = parseFile("example.txt")

func TestToDecimal(t *testing.T) {
	tcs := []struct {
		snafu string
		want  int
	}{
		{"1=-0-2", 1747},
		{"12111", 906},
		{"1121-1110-1=0", 314159265},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, toDecimal(tc.snafu))
		})
	}
}

func TestToSnafu(t *testing.T) {
	tcs := []struct {
		decimal int
		want    string
	}{
		{1747, "1=-0-2"},
		{3, "1="},
		{4, "1-"},
		{5, "10"},
		{6, "11"},
		{7, "12"},
		{8, "2="},
		{9, "2-"},
		{10, "20"},
		{15, "1=0"},
		{20, "1-0"},
		{2022, "1=11-2"},
		{906, "12111"},
		{314159265, "1121-1110-1=0"},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, toSnafu(tc.decimal))
		})
	}
}

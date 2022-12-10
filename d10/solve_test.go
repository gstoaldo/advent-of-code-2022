package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example1 = parseFile("example1.txt")
var example2 = parseFile("example2.txt")

func TestRegisterValueAtCycle(t *testing.T) {
	tcs := []struct {
		instructions inputType
		ncycle       int
		want         int
	}{
		{example1, 1, 1},
		{example1, 2, 1},
		{example1, 3, 1},
		{example1, 4, 4},
		{example1, 5, 4},
		{example1, 6, -1},
		//
		{example2, 20, 21},
		{example2, 60, 19},
		{example2, 100, 18},
		{example2, 140, 21},
		{example2, 180, 16},
		{example2, 220, 18},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, registerValueAtCycle(tc.instructions, tc.ncycle))
		})
	}
}

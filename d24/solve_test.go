package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example = parseFile("example2.txt")

func TestBlizzardEq(t *testing.T) {
	blizzards := append([]blizzardT{}, example.blizzards...)

	tcs := []struct {
		blizzardId int
		t          int
		want       positionT
	}{
		{3, 0, positionT{1, 5}},
		{3, 1, positionT{4, 5}},
		{3, 4, positionT{1, 5}},
		{3, 5, positionT{4, 5}},
		{3, 8, positionT{1, 5}},
		{3, 9, positionT{4, 5}},
		{0, 0, positionT{1, 1}},
		{0, 1, positionT{1, 2}},
		{0, 6, positionT{1, 1}},
		{0, 7, positionT{1, 2}},
		{0, 12, positionT{1, 1}},
		{0, 18, positionT{1, 1}},
		{4, 0, positionT{1, 6}},
		{4, 6, positionT{1, 6}},
		{4, 12, positionT{1, 6}},
		{4, 18, positionT{1, 6}},
	}

	maxI := example.maxI
	maxJ := example.maxJ

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			blizzard := blizzards[tc.blizzardId]
			got := blizzardEq(blizzard, tc.t, maxI, maxJ)
			utils.Assert(t, tc.want, got)
		})
	}
}

package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetElfProposal(t *testing.T) {
	elvesMap := elvesMapT(parseFile("example2.txt"))
	elvesMap[positionT{100, 100}] = 1

	directions := []directionT{N, S, W, E}
	tcs := []struct {
		elf  positionT
		want positionT
	}{
		{positionT{1, 2}, positionT{0, 2}},
		{positionT{1, 3}, positionT{0, 3}},
		{positionT{2, 2}, positionT{3, 2}},
		{positionT{4, 2}, positionT{3, 2}},
		{positionT{100, 100}, positionT{100, 100}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := getElfProposal(elvesMap, tc.elf, directions)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestGetNextElvesMap(t *testing.T) {
	elvesMap := elvesMapT(parseFile("example2.txt"))
	directions := []directionT{N, S, W, E}
	want := elvesMapT{
		{0, 2}: 1,
		{0, 3}: 1,
		{2, 2}: 1,
		{3, 3}: 1,
		{4, 2}: 1,
	}

	got := getNextElvesMap(elvesMap, directions)
	assert.Equal(t, want, got)
}

// func TestGetFinalElvesMapSmallExample(t *testing.T) {
// 	elvesMap := elvesMapT(example2)
// 	want := elvesMapT{
// 		{0, 2}: 1,
// 		{1, 4}: 1,
// 		{2, 0}: 1,
// 		{3, 4}: 1,
// 		{5, 2}: 1,
// 	}

// 	finalMap := getFinalElvesMap(elvesMap, 10)
// 	assert.Equal(t, want, finalMap)

// 	count := countFreeTiles(finalMap)
// 	assert.Equal(t, 25, count)
// }

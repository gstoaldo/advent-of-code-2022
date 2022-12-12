package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
)

var example1 = parseFile("example1.txt")

func TestFindNeighbours(t *testing.T) {
	tcs := []struct {
		grid [][]string
		p    position
		want []position
	}{
		{example1.grid, position{0, 0}, []position{{0, 1}, {1, 0}}},
		{example1.grid, position{1, 0}, []position{{1, 1}, {2, 0}, {0, 0}}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			neighbors := findNeighbors(tc.grid, tc.p)
			assert.Equal(t, tc.want, neighbors)
		})
	}
}

func TestBFS(t *testing.T) {
	got := bfs(example1.grid, example1.start, example1.end)
	utils.Assert(t, 31, got)
}

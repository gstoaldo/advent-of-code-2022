package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = parseFile("example.txt")

func TestRunMove(t *testing.T) {
	tcs := []struct {
		stacks [][]string
		move   []int
		want   [][]string
	}{
		{
			example.stacks, example.moves[0], [][]string{{"Z", "N", "D"}, {"M", "C"}, {"P"}},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := runMove(tc.stacks, tc.move, 1)

			assert.Equal(t, tc.want, got)
		})
	}
}

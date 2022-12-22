package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example = parseFile("example.txt")

func TestWrap(t *testing.T) {
	tcs := []struct {
		board        boardT
		currPosition pointT
		vector       vectorT
		want         pointT
	}{
		{example.board, pointT{12, 7}, DIRECTIONS[0], pointT{1, 7}},
		{example.board, pointT{6, 8}, DIRECTIONS[1], pointT{6, 5}},
		{example.board, pointT{12, 12}, DIRECTIONS[1], pointT{12, 12}},
		{example.board, pointT{1, 5}, DIRECTIONS[2], pointT{1, 5}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, wrap(tc.board, tc.currPosition, tc.vector))
		})
	}
}

func TestMove(t *testing.T) {
	tcs := []struct {
		board        boardT
		currPosition pointT
		vector       vectorT
		nTiles       int
		want         pointT
	}{
		{example.board, pointT{9, 1}, DIRECTIONS[0], 2, pointT{11, 1}},
		{example.board, pointT{9, 1}, DIRECTIONS[0], 20, pointT{11, 1}},
		{example.board, pointT{11, 1}, DIRECTIONS[2], 20, pointT{9, 1}},
		{example.board, pointT{11, 1}, DIRECTIONS[2], 1, pointT{10, 1}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, move(tc.board, tc.currPosition, tc.vector, tc.nTiles))
		})
	}
}

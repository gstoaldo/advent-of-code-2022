package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example = parseFile("example.txt")

func TestWrap2D(t *testing.T) {
	tcs := []struct {
		board        boardT
		currPosition pointT
		facing       int
		want         pointT
	}{
		{example.board, pointT{12, 7}, 0, pointT{1, 7}},
		{example.board, pointT{6, 8}, 1, pointT{6, 5}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got, _ := wrap2D(tc.board, tc.currPosition, tc.facing)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestMove(t *testing.T) {
	tcs := []struct {
		board        boardT
		currPosition pointT
		facing       int
		nTiles       int
		want         pointT
	}{
		{example.board, pointT{9, 1}, 0, 2, pointT{11, 1}},
		{example.board, pointT{9, 1}, 0, 20, pointT{11, 1}},
		{example.board, pointT{11, 1}, 2, 20, pointT{9, 1}},
		{example.board, pointT{11, 1}, 2, 1, pointT{10, 1}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got, _ := move(tc.board, tc.currPosition, tc.facing, tc.nTiles, wrap2D)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestConvertCoords(t *testing.T) {
	tcs := []struct {
		p0     pointT
		face   int
		facing int
		want   pointT
	}{
		{pointT{7, 5}, 3, 3, pointT{9, 3}},
		{pointT{9, 3}, 1, 2, pointT{7, 5}},
		{pointT{12, 6}, 4, 0, pointT{15, 9}},
		{pointT{15, 9}, 6, 3, pointT{12, 6}},
		{pointT{11, 12}, 5, 1, pointT{2, 8}},
		{pointT{2, 8}, 2, 1, pointT{11, 12}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got, _ := convertCoords(tc.p0, tc.face, tc.facing)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestCurrentFace(t *testing.T) {
	tcs := []struct {
		board     boardT
		currPoint pointT
		want      int
	}{
		{example.board, pointT{12, 1}, 1},
		{example.board, pointT{1, 5}, 2},
		{example.board, pointT{2, 6}, 2},
		{example.board, pointT{4, 8}, 2},
		{example.board, pointT{9, 6}, 4},
		{example.board, pointT{16, 12}, 6},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, currentFace(tc.board, tc.currPoint, 4))
		})
	}
}

func TestConvertCoordsInput(t *testing.T) {
	tcs := []struct {
		p0     pointT
		face   int
		facing int
		want   pointT
	}{
		{pointT{51, 1}, 1, 3, pointT{1, 151}},
		{pointT{102, 1}, 2, 3, pointT{2, 200}},
		{pointT{150, 50}, 2, 0, pointT{100, 101}},
		{pointT{100, 101}, 5, 0, pointT{150, 50}},
	}

	CUBE_FACE_MAP = inputFaceMap

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got, _ := convertCoords(tc.p0, tc.face, tc.facing)
			utils.Assert(t, tc.want, got)
		})
	}
}

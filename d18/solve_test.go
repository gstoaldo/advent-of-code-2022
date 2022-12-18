package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example1 = parseFile("example1.txt")
var example2 = parseFile("example2.txt")

func TestIsTouchingXY(t *testing.T) {
	tcs := []struct {
		cube1, cube2 cubeT
		want         bool
	}{
		{cubeT{0, 0, 0}, cubeT{0, 1, 0}, true},
		{cubeT{0, 0, 0}, cubeT{1, 0, 0}, true},
		{cubeT{0, 0, 0}, cubeT{1, 1, 0}, false},
		{cubeT{0, 0, 0}, cubeT{0, 2, 0}, false},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, isTouchingXY(tc.cube1, tc.cube2))
		})
	}
}

func TestCountTouchingSidesXY(t *testing.T) {
	tcs := []struct {
		cubesXY []cubeT
		want    int
	}{
		{filterByZ(example1, 1), 0},
		{filterByZ(example1, 2), 4},
		{filterByZ(example1, 5), 0},
		{[]cubeT{{0, 0, 0}, {0, 1, 0}, {0, 2, 0}, {0, 3, 0}, {1, 3, 0}}, 4},
		{[]cubeT{{0, 0, 0}, {0, 1, 0}, {0, 2, 0}, {0, 3, 0}, {1, 3, 0}}, 4},
		{filterByZ(example2, 0), 40},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, countTouchingSidesXY(tc.cubesXY))
		})
	}
}

func TestCountTouchingSidesZ(t *testing.T) {
	tcs := []struct {
		cubesZ1, cubesZ2 []cubeT
		want             int
	}{
		{filterByZ(example1, 1), filterByZ(example1, 2), 1},
		{filterByZ(example1, 2), filterByZ(example1, 3), 1},
		{filterByZ(example1, 4), filterByZ(example1, 5), 0},
		{filterByZ(example1, 1), []cubeT{}, 0},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, countTouchingSidesZ(tc.cubesZ1, tc.cubesZ2))
		})
	}
}

func TestCountFreeSides3D(t *testing.T) {
	utils.Assert(t, 64, countFreeSides3D(example1))
}

func TestGetAllAirPockets3D(t *testing.T) {
	tcs := []struct {
		cubes []cubeT
		want  int
	}{
		{example1, 1},
		{example2, 1},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := getAllAirPockets3D(tc.cubes)
			utils.Assert(t, tc.want, len(got))
		})
	}
}

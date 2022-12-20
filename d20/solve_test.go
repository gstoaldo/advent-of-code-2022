package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
)

var example = parseFile("example.txt")

func TestCalcNewId(t *testing.T) {
	tcs := []struct {
		value, oldId, length, want int
	}{
		{1, 0, 7, 1},
		{2, 0, 7, 2},
		{-3, 1, 7, 4},
		{3, 2, 7, 5},
		{-2, 2, 7, 6},
		{0, 3, 7, 3},
		{4, 5, 7, 3},
		{8, 1, 3, 1},
		{6, 0, 7, 0},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			utils.Assert(t, tc.want, calcNewId(tc.value, tc.oldId, tc.length))
		})
	}
}

func TestMove(t *testing.T) {
	tcs := []struct {
		value       int
		id          int
		arrangement []int
		want        []int
	}{
		{1, 0, []int{1, 2, -3, 3, -2, 0, 4}, []int{2, 1, -3, 3, -2, 0, 4}},
		{2, 0, []int{2, 1, -3, 3, -2, 0, 4}, []int{1, -3, 2, 3, -2, 0, 4}},
		{-3, 1, []int{1, -3, 2, 3, -2, 0, 4}, []int{1, 2, 3, -2, -3, 0, 4}},
		{3, 2, []int{1, 2, 3, -2, -3, 0, 4}, []int{1, 2, -2, -3, 0, 3, 4}},
		{-2, 2, []int{1, 2, -2, -3, 0, 3, 4}, []int{1, 2, -3, 0, 3, 4, -2}},
		{0, 3, []int{1, 2, -3, 0, 3, 4, -2}, []int{1, 2, -3, 0, 3, 4, -2}},
		{4, 5, []int{1, 2, -3, 0, 3, 4, -2}, []int{1, 2, -3, 4, 0, 3, -2}},
		{1, 0, []int{1, 2, 1}, []int{2, 1, 1}},
		{1, 2, []int{1, 2, 1}, []int{1, 1, 2}},
		{1, 2, []int{1, 2, 1, 4}, []int{1, 2, 4, 1}},
		{1, 3, []int{4, 5, 6, 1, 7, 8, 9}, []int{4, 5, 6, 7, 1, 8, 9}},
		{-2, 1, []int{4, -2, 5, 6, 7, 8, 9}, []int{4, 5, 6, 7, 8, -2, 9}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			inputMap := getInputMap(tc.arrangement)
			move(numberT{tc.value, tc.id}, inputMap)

			for number, id := range inputMap {
				utils.Assert(t, tc.want[id], number.value)
			}
		})
	}
}

func TestMix(t *testing.T) {
	want := []int{1, 2, -3, 4, 0, 3, -2}
	got := mix(example, 1)
	assert.Equal(t, want, got)
}

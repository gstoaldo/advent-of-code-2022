package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetXIntervalAtDistance(t *testing.T) {
	tcs := []struct {
		sensor pointT
		dist   int
		y      int
		want   intervalT
	}{
		{pointT{8, 7}, 9, 5, intervalT{1, 15}},
		{pointT{14, 3}, 1, 3, intervalT{13, 15}},
		{pointT{0, 0}, 1, 10, intervalT{}},
		{pointT{8, 7}, 9, -2, intervalT{8, 8}},
		{pointT{8, 7}, 9, -1, intervalT{7, 9}},
		{pointT{8, 7}, 9, 0, intervalT{6, 10}},
		{pointT{8, 7}, 9, 7, intervalT{-1, 17}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := getXIntervalAtDistance(tc.sensor, tc.dist, tc.y)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestMergeIntervals(t *testing.T) {
	tcs := []struct {
		intervals []intervalT
		want      []intervalT
	}{
		{[]intervalT{{1, 13}, {15, 16}, {13, 15}}, []intervalT{{1, 16}}},
		{[]intervalT{{1, 13}, {15, 16}}, []intervalT{{1, 13}, {15, 16}}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.want, mergeIntervals(tc.intervals))
		})
	}
}

package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
)

var example1 = parseFile("example1.txt")

func TestGetPathOptimalValves(t *testing.T) {
	tcs := []struct {
		graph inputT
		path  []string
		want  []int
	}{
		{
			example1,
			[]string{"AA", "DD", "CC", "BB", "AA", "II", "JJ", "II", "AA", "DD", "EE", "FF", "GG", "HH", "GG", "FF", "EE", "DD", "CC"},
			[]int{1, 3, 6, 13, 16, 18},
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := getPathOptimalValves(tc.graph, tc.path)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPressure(t *testing.T) {
	tcs := []struct {
		graph         inputT
		path          []string
		openValvesIds []int
		want          int
	}{
		{
			example1,
			[]string{"AA", "DD", "CC", "BB", "AA", "II", "JJ", "II", "AA", "DD", "EE", "FF", "GG", "HH", "GG", "FF", "EE", "DD", "CC"},
			[]int{1, 3, 6, 13, 16, 18},
			1651,
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := pressure(tc.graph, tc.path, tc.openValvesIds)
			utils.Assert(t, tc.want, got)
		})
	}
}

package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
)

var example = parseFile("example.txt")

func TestFindDirectoryEndIndex(t *testing.T) {
	tcs := []struct {
		output     inputType
		startIndex int
		want       int
	}{
		{example, 6, 16},
		{example, 12, 15},
		{example, 0, len(example) - 1},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := findDirectoryEndIndex(tc.output, tc.startIndex)
			utils.Assert(t, tc.want, got)
		})
	}

}

func TestGetDirectoryTotalSize(t *testing.T) {
	tcs := []struct {
		output     inputType
		startIndex int
		want       int
	}{
		{example, 6, 94853},
		{example, 12, 584},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			endIndex := findDirectoryEndIndex(tc.output, tc.startIndex)
			got := getDirectoryTotalSize(tc.output, tc.startIndex, endIndex)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestGetAllDirectorySize(t *testing.T) {
	want := []int{48381165, 94853, 584, 24933642}
	assert.Equal(t, want, getAllDirectorySize(example))
}

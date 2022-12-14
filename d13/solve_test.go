package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestFindStartIndex(t *testing.T) {
	tcs := []struct {
		s      string
		target []int
		want   int
	}{
		{"[[1],[2,3,4]]", []int{0}, 0},
		{"[[1],[2,3,4]]", []int{0, 0}, 1},
		{"[[1],[2,3,4]]", []int{1}, 4},
		{"[[1],[2,3,4]]", []int{1, 0}, 5},
		{"[[1],[2,3,4]]", []int{2}, -1},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := findStartIndex(tc.s, tc.target)
			utils.Assert(t, tc.want, got)
		})
	}
}
func TestFindCloseIndex(t *testing.T) {
	tcs := []struct {
		s          string
		startIndex int
		want       int
	}{
		{"[[1],[2,3,4]]", 0, 4},
		{"[[1],[2,3,4]]", 1, 3},
		{"[[100],[2,3,4]]", 1, 5},
		{"[[1],[2,3,4]]", 5, 7},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := findCloseIndex(tc.s, tc.startIndex)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestFindElementByLevel(t *testing.T) {
	tcs := []struct {
		s      string
		target []int
		want   string
	}{
		{"[[1],[2,3,4]]", []int{0}, "[1]"},
		{"[[1],[2,3,4]]", []int{0, 0}, "1"},
		{"[[1],[2,3,4]]", []int{1}, "[2,3,4]"},
		{"[[1],[2,3,4]]", []int{1, 0}, "2"},
		{"[[1],[2,3,4]]", []int{1, 1}, "3"},
		{"[[1],[2,3,4]]", []int{1, 2}, "4"},
		{"[[1],[2,3,4]]", []int{2}, ""},
		{"[7,7,7,7]", []int{0}, "7"},
		{"[[8,7,6]]", []int{0}, "[8,7,6]"},
		{"[[8,7,6]]", []int{0, 0}, "8"},
		{"[[[]]]", []int{0}, "[[]]"},
		{"[[[]]]", []int{0, 0}, "[]"},
		{"[[[]]]", []int{0, 0, 0}, ""},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", []int{1, 0}, "2"},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", []int{1, 1}, "[3,[4,[5,6,7]]]"},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := findElementByLevel(tc.s, tc.target)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestPairIsOrdered(t *testing.T) {
	tcs := []struct {
		p1   string
		p2   string
		want bool
	}{
		{"[1,1,3,1,1]", "[1,1,5,1,1]", true},
		{"[[1],[2,3,4]]", "[[1],4]", true},
		{"[9]", "[[8,7,6]]", false},
		{"[[4,4],4,4]", "[[4,4],4,4,4]", true},
		{"[7,7,7,7]", "[7,7,7]", false},
		{"[]", "[3]", true},
		{"[[[]]]", "[[]]", false},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]", false},
		{"[[2,7],[1,[4,[9,9,6]],0]]", "[[[[5,2]],10,5]]", true},
		{"[[2,7],[1,[4,[9,9,6]],0]]", "[[[[1,2]],10,5]]", false},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {

			_, got := pairIsOrdered(tc.p1, tc.p2)
			assert.Equal(t, tc.want, got)

		})
	}
}

func TestSortPackages(t *testing.T) {
	tcs := []struct {
		packages []string
		want     []string
	}{
		{[]string{"[[1],4]", "[[2]]"}, []string{"[[1],4]", "[[2]]"}},
		{[]string{"[[2]]", "[[1],4]"}, []string{"[[1],4]", "[[2]]"}},
		{[]string{"[[[]]]", "[[]]", "[]"}, []string{"[]", "[[]]", "[[[]]]"}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			sortPackages(tc.packages)
			assert.Equal(t, tc.want, tc.packages)
		})
	}
}

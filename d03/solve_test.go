package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example = parseFile("example.txt")

func TestGetItemPriority(t *testing.T) {
	tcs := []struct {
		item string
		want int
	}{
		{"p", 16},
		{"L", 38},
		{"P", 42},
		{"v", 22},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := getItemPriority(tc.item)
			utils.Assert(t, tc.want, got)
		})
	}
}

func TestGetRepeatedItem(t *testing.T) {
	tcs := []struct {
		bag  string
		want string
	}{
		{example[0], "p"},
		{example[1], "L"},
		{example[2], "P"},
		{example[3], "v"},
		{example[4], "t"},
		{example[5], "s"},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := getRepeatedItem(tc.bag)
			utils.Assert(t, tc.want, got)
		})
	}

}

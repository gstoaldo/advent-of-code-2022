package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

var example1 = parseFile("example1.txt")

func TestGetNextPoint(t *testing.T) {
	var rocks = getRocks(example1)

	tcs := []struct {
		rocks blockedT
		sands blockedT
		p0    pointT
		want  pointT
	}{
		{rocks, blockedT{}, pointT{500, 0}, pointT{500, 1}},
		{rocks, blockedT{}, pointT{500, 8}, pointT{500, 8}},
		{rocks, blockedT{pointT{500, 8}: true, pointT{499, 8}: true}, pointT{500, 7}, pointT{501, 8}},
		{rocks, blockedT{pointT{500, 8}: true, pointT{499, 8}: true, pointT{501, 8}: true}, pointT{500, 7}, pointT{500, 7}},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			got := getNextPoint(tc.p0, tc.rocks, tc.sands)
			utils.Assert(t, tc.want, got)
		})
	}
}

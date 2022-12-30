package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = parseFile("example.txt")

func TestBuildNow(t *testing.T) {

	state := stateT{
		resources: []int{2, 0, 0, 0},
		robots:    []int{1, 0, 0, 0},
		minute:    2,
	}
	bp := example[0]

	nextState := buildNow(state, bp[1], 1)

	assert.Equal(t, 3, nextState.minute)
	assert.Equal(t, []int{1, 0, 0, 0}, nextState.resources)
	assert.Equal(t, []int{1, 1, 0, 0}, nextState.robots)
}

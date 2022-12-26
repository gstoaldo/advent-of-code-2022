package main

import (
	"io/ioutil"
	"strings"
)

var charToVelocity = map[rune]velocityT{
	'v': {1, 0},
	'^': {-1, 0},
	'>': {0, 1},
	'<': {0, -1},
}

func parseFile(path string) inputT {
	file, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(file), "\n")
	var start, end positionT

	maxI := len(lines) - 1
	maxJ := len(lines[0]) - 1
	blizzards := []blizzardT{}

	for i, line := range lines {
		for j, char := range line {
			if i == 0 && char == '.' {
				start = positionT{i, j}
			}

			if i == maxI && char == '.' {
				end = positionT{i, j}
			}

			if char != '.' && char != '#' {
				blizzards = append(blizzards, blizzardT{positionT{i, j}, charToVelocity[char]})
			}
		}
	}

	return inputT{
		start:     start,
		end:       end,
		blizzards: blizzards,
		maxI:      maxI,
		maxJ:      maxJ,
	}
}

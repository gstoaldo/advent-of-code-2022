package main

import (
	"io/ioutil"
	"strings"
)

func parser(path string) inputT {
	file, _ := ioutil.ReadFile(path)

	input := inputT{}
	lines := strings.Split(string(file), "\n")

	for row, line := range lines {
		for col, r := range line {
			if r == '#' {
				input[positionT{row, col}] = 1
			}
		}
	}

	return input
}

func parseFile(path string) inputT {
	return parser(path)
}

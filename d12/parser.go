package main

import (
	"io/ioutil"
	"strings"
)

func parser(path string) inputType {
	file, _ := ioutil.ReadFile(path)

	lines := strings.Split(string(file), "\n")

	grid := [][]string{}
	start, end := position{}, position{}

	for i, line := range lines {
		row := []string{}

		for j, r := range line {
			s := string(r)
			if s == "S" {
				start = position{i, j}
				s = "a"
			}

			if s == "E" {
				end = position{i, j}
				s = "z"
			}

			row = append(row, s)
		}
		grid = append(grid, row)
	}

	return inputType{grid, start, end}
}

func parseFile(path string) inputType {
	return parser(path)
}

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func parser(path string) inputT {
	file, _ := ioutil.ReadFile(path)

	lines := strings.Split(string(file), "\n")

	input := inputT{}

	for _, line := range lines {
		linePoints := []pointT{}
		points := strings.Split(line, " -> ")

		for _, point := range points {
			coords := strings.Split(point, ",")

			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])

			linePoints = append(linePoints, pointT{x, y})
		}

		input = append(input, linePoints)
	}

	return input
}

func parseFile(path string) inputT {
	return parser(path)
}

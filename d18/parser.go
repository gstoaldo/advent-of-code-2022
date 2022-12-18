package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func parser(path string) inputT {
	file, _ := ioutil.ReadFile(path)

	lines := strings.Split(string(file), "\n")

	cubes := inputT{}

	for _, line := range lines {
		coords := strings.Split(line, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		cubes = append(cubes, cubeT{x, y, z})
	}

	return cubes
}

func parseFile(path string) inputT {
	return parser(path)
}

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
		x, _ := strconv.Atoi(line)
		input = append(input, x)
	}

	return input
}

func parseFile(path string) inputT {
	return parser(path)
}

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func parser(path string) inputType {
	file, _ := ioutil.ReadFile(path)

	lines := strings.Split(string(file), "\n")

	moves := inputType{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		direction := parts[0]
		value, _ := strconv.Atoi(parts[1])
		move := stepType{direction, value}
		moves = append(moves, move)
	}

	return moves
}

func parseFile(path string) inputType {
	return parser(path)
}

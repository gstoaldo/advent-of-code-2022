package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func parser(path string) inputType {
	file, _ := ioutil.ReadFile(path)

	lines := strings.Split(string(file), "\n")

	instructions := inputType{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		name := parts[0]
		value := 0

		if len(parts) > 1 {
			value, _ = strconv.Atoi(parts[1])
		}

		instructions = append(instructions, instruction{name, value})
	}

	return instructions
}

func parseFile(path string) inputType {
	return parser(path)
}

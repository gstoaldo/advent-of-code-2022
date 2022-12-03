package main

import (
	"bufio"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

func parser(scanner *bufio.Scanner) interface{} {
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseFile(path string) []string {
	input := utils.ParseFile(path, parser)
	lines := input.([]string)

	return lines
}

package main

import (
	"bufio"
	"strings"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

func parser(scanner *bufio.Scanner) interface{} {
	rounds := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		rounds = append(rounds, strings.Split(line, " "))
	}

	return rounds
}

func parseFile(path string) [][]string {
	input := utils.ParseFile(path, parser)
	rounds := input.([][]string)

	return rounds
}

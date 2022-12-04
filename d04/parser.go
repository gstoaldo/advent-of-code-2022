package main

import (
	"bufio"
	"regexp"
	"strconv"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

func parser(scanner *bufio.Scanner) interface{} {
	lines := [][]int{}

	re := regexp.MustCompile(`\d{1,}`)

	for scanner.Scan() {
		line := scanner.Text()
		idsString := re.FindAllString(line, -1)

		lineInt := []int{}
		for _, id := range idsString {
			idInt, _ := strconv.Atoi(id)
			lineInt = append(lineInt, idInt)
		}

		lines = append(lines, lineInt)
	}

	return lines
}

func parseFile(path string) [][]int {
	input := utils.ParseFile(path, parser)
	rounds := input.([][]int)

	return rounds
}

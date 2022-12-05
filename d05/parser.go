package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type inputType struct {
	stacks [][]string
	moves  [][]int
}

func parseFirstPart(lines []string) [][]string {
	nstacks := (len(lines[0])-3)/4 + 1
	stacks := make([][]string, nstacks)
	re := regexp.MustCompile(`[A-Z]`)

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]

		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			loc := match[0]
			crate := string(line[loc])
			stackId := (loc - 1) / 4

			stacks[stackId] = append([]string{crate}, stacks[stackId]...)
		}
	}

	return stacks
}

func parseSecondPart(lines []string) [][]int {
	re := regexp.MustCompile(`\d{1,}`)
	moves := [][]int{}

	for _, line := range lines {
		steps := re.FindAllString(line, -1)

		stepsInt := []int{}

		for i, step := range steps {
			delta := 0

			if i > 0 {
				delta = 1
			}

			stepInt, _ := strconv.Atoi(step)
			stepsInt = append(stepsInt, stepInt-delta)
		}

		moves = append(moves, stepsInt)
	}

	return moves
}

func parser(path string) inputType {
	file, _ := ioutil.ReadFile(path)

	fileContent := string(file)

	parts := strings.Split(fileContent, "\n\n")

	stacks := parseFirstPart(strings.Split(parts[0], "\n"))

	moves := parseSecondPart(strings.Split(parts[1], "\n"))

	return inputType{
		stacks: stacks,
		moves:  moves,
	}
}

func parseFile(path string) inputType {
	return parser(path)
}

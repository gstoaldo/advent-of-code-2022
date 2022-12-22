package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func parser(path string) inputT {
	file, _ := ioutil.ReadFile(path)

	re := regexp.MustCompile(`\d+|[LR]`)

	parts := strings.Split(string(file), "\n\n")
	boardRows := strings.Split(parts[0], "\n")

	board := boardT{}
	start := pointT{}
	startFound := false

	for y, row := range boardRows {
		for x, r := range row {
			if r == ' ' {
				continue
			}

			isOpen := r == '.'

			if !startFound && isOpen {
				start = pointT{x + 1, y + 1}
				startFound = true
			}

			board[pointT{x + 1, y + 1}] = isOpen
		}
	}

	commands := re.FindAllString(parts[1], -1)

	return inputT{
		board:    board,
		commands: commands,
		start:    start,
	}

}

func parseFile(path string) inputT {
	return parser(path)
}

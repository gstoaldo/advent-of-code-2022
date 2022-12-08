package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func parser(path string) inputType {
	file, _ := ioutil.ReadFile(path)

	lines := strings.Split(string(file), "\n")

	rows := inputType{}

	for _, line := range lines {
		row := []int{}

		for _, n := range line {
			nint, _ := strconv.Atoi(string(n))
			row = append(row, nint)
		}
		rows = append(rows, row)
	}

	return rows
}

func parseFile(path string) inputType {
	return parser(path)
}

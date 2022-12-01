package main

import (
	"bufio"
	"strconv"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

func parser(scanner *bufio.Scanner) interface{} {
	elf := Elf{}
	elves := []Elf{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elves = append(elves, elf)
			elf = Elf{}
		} else {
			cal, _ := strconv.Atoi(line)
			elf = append(elf, cal)
		}
	}

	elves = append(elves, elf)

	return elves
}

func parseFile(path string) []Elf {
	input := utils.ParseFile(path, parser)
	elves := input.([]Elf)

	return elves
}

package main

import (
	"fmt"
)

type Elf []int

func getElfMaxCalories(elves []Elf) int {
	var max int
	for _, elf := range elves {
		var totalCal int
		for _, cal := range elf {
			totalCal += cal
		}

		if totalCal > max {
			max = totalCal
		}
	}

	return max
}

func part1(elves []Elf) {
	part1 := getElfMaxCalories(elves)
	fmt.Println("part 1:", part1)
}

func main() {
	elves := parseFile("input.txt")

	part1(elves)
}

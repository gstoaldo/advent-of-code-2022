package main

import (
	"fmt"
)

type inputType []instruction

type instruction struct {
	name  string
	value int
}

func registerValueAtCycle(instructions inputType, ncycle int) int {
	reg := 1
	currCycle := 1
	valueToAdd := 0

	for _, instruction := range instructions {
		if currCycle >= ncycle {
			break
		}

		reg += valueToAdd
		valueToAdd = 0

		if instruction.name == "noop" {
			currCycle++
		}

		if instruction.name == "addx" {
			valueToAdd = instruction.value
			currCycle += 2
		}
	}

	if currCycle == ncycle {
		reg += valueToAdd
	}

	return reg
}

func calcSumSignalStrenth(instructions inputType) int {
	sum := 0

	for ncycle := 20; ncycle <= 220; ncycle += 40 {
		sum += ncycle * registerValueAtCycle(instructions, ncycle)
	}

	return sum
}

func mapCycleToGrid(ncycle int) (int, int) {
	return (ncycle - 1) / 40, (ncycle - 1) % 40
}

func renderGrid(grid [6][40]bool) {
	for _, row := range grid {
		for _, isLit := range row {
			if isLit {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func render(instructions inputType) {
	grid := [6][40]bool{}

	for ncycle := 1; ncycle <= 240; ncycle++ {
		position := registerValueAtCycle(instructions, ncycle)
		i, j := mapCycleToGrid(ncycle)

		grid[i][j] = j >= position-1 && j <= position+1
	}

	renderGrid(grid)
}

func part1(input inputType) {
	answer := calcSumSignalStrenth(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	render(input)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}

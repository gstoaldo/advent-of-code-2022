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

func part1(input inputType) {
	answer := calcSumSignalStrenth(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	answer := ""
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}

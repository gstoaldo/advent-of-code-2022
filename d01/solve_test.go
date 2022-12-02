package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

func TestGetElfMaxCalories(t *testing.T) {
	input := utils.ParseFile("example.txt", parser)
	elves := input.([]Elf)

	got := getElfMaxCalories(elves)
	utils.Assert(t, 24000, got)
}

func TestGetTop3CaloriesSum(t *testing.T) {
	input := utils.ParseFile("example.txt", parser)
	elves := input.([]Elf)

	got := getTop3CaloriesSum(elves)
	utils.Assert(t, 45000, got)
}

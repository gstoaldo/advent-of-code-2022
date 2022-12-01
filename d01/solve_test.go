package main

import (
	"testing"

	"github.com/gstoaldo/advent-of-code-2022/utils"
)

func TestGetElfMaxCalories(t *testing.T) {
	input := utils.ParseFile("example.txt", parser)
	elves := input.([]Elf)

	want := 24000
	got := getElfMaxCalories(elves)

	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestGetTop3CaloriesSum(t *testing.T) {
	input := utils.ParseFile("example.txt", parser)
	elves := input.([]Elf)

	want := 45000
	got := getTop3CaloriesSum(elves)

	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

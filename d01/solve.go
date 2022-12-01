package main

import (
	"fmt"
	"sort"
)

type Elf []int

func sumArray(array []int) int {
	var sum int
	for _, v := range array {
		sum += v
	}
	return sum
}

func getTotalCaloriesByElfSorted(elves []Elf) []int {
	calByElf := []int{}

	for _, elf := range elves {
		calByElf = append(calByElf, sumArray(elf))
	}

	sort.Slice(calByElf, func(i, j int) bool {
		return calByElf[i] > calByElf[j]
	})

	return calByElf
}

func getElfMaxCalories(elves []Elf) int {
	calByElf := getTotalCaloriesByElfSorted(elves)

	return calByElf[0]
}

func getTop3CaloriesSum(elves []Elf) int {
	caloriesByElf := getTotalCaloriesByElfSorted(elves)

	return sumArray(caloriesByElf[0:3])

}

func part1(elves []Elf) {
	answer := getElfMaxCalories(elves)
	fmt.Println("part 1:", answer)
}

func part2(elves []Elf) {
	answer := getTop3CaloriesSum(elves)
	fmt.Println("part 2:", answer)
}

func main() {
	elves := parseFile("input.txt")

	part1(elves)
	part2(elves)
}

package main

import (
	"fmt"
	"sort"
)

type inputType struct {
	monkeys []monkey
	mmc     int
}

type monkey struct {
	items               []int
	opsFunc             func(int) int
	nextFunc            func(int) int
	inspectedItemsCount int
}

func runRounds(monkeys []monkey, nrounds int, relief func(int) int) {
	for n := 0; n < nrounds; n++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				nextWorryLevel := relief(monkey.opsFunc(item))
				nextMonkey := monkey.nextFunc(nextWorryLevel)

				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, nextWorryLevel)
			}
			monkeys[i].inspectedItemsCount += len(monkey.items)
			monkeys[i].items = []int{}
		}
	}
}

func calcMonkeyBusinessLevel(monkeys []monkey) int {
	inspectedItemsCount := []int{}

	for _, m := range monkeys {
		inspectedItemsCount = append(inspectedItemsCount, m.inspectedItemsCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspectedItemsCount)))

	return inspectedItemsCount[0] * inspectedItemsCount[1]
}

func reliefPT1(v int) int {
	return v / 3
}

func getReliefPT2(mmc int) func(int) int {
	return func(v int) int {
		return v % mmc
	}
}

func part1(input inputType) {
	cp := make([]monkey, len(input.monkeys))
	copy(cp, input.monkeys)

	relief := reliefPT1

	runRounds(cp, 20, relief)
	answer := calcMonkeyBusinessLevel(cp)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	cp := make([]monkey, len(input.monkeys))
	copy(cp, input.monkeys)

	relief := getReliefPT2(input.mmc)

	runRounds(cp, 10000, relief)
	answer := calcMonkeyBusinessLevel(cp)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}

package main

import (
	"fmt"
	"sort"
)

type inputType []monkey

type monkey struct {
	items               []int
	opsFunc             func(int) int
	nextFunc            func(int) int
	inspectedItemsCount int
}

func runRounds(monkeys inputType, nrounds int) {
	for n := 0; n < nrounds; n++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				nextWorryLevel := monkey.opsFunc(item) / 3
				nextMonkey := monkey.nextFunc(nextWorryLevel)

				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, nextWorryLevel)
			}
			monkeys[i].inspectedItemsCount += len(monkey.items)
			monkeys[i].items = []int{}
		}
	}
}

func calcMonkeyBusinessLevel(monkeys inputType) int {
	inspectedItemsCount := []int{}

	for _, m := range monkeys {
		inspectedItemsCount = append(inspectedItemsCount, m.inspectedItemsCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspectedItemsCount)))

	return inspectedItemsCount[0] * inspectedItemsCount[1]
}

func part1(input inputType) {
	cp := make(inputType, len(input))
	copy(cp, input)

	runRounds(cp, 20)
	answer := calcMonkeyBusinessLevel(cp)
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

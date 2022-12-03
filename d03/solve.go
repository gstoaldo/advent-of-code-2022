package main

import (
	"fmt"
	"unicode"
)

func getItemPriority(item string) int {
	base := int('a')
	itemRune := rune(item[0])

	if !unicode.IsLower(itemRune) {
		base = int('A') - 26
	}

	return int(itemRune) - base + 1
}

func getRepeatedItem(bag string) string {
	leftComp := bag[:len(bag)/2]
	rightComp := bag[len(bag)/2:]

	leftCompSet := map[rune]bool{}

	for _, item := range leftComp {
		leftCompSet[item] = true
	}

	for _, item := range rightComp {
		if _, ok := leftCompSet[item]; ok {
			return string(item)
		}
	}

	return ""
}

func getPrioritySum(bags []string) int {
	sum := 0

	for _, bag := range bags {
		repeatedItem := getRepeatedItem(bag)
		sum += getItemPriority(repeatedItem)
	}

	return sum
}

func part1(input []string) {
	answer := getPrioritySum(input)
	fmt.Println("part 1:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
}

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

func getCommomItems(itemsA string, itemsB string) string {
	commomItems := ""

	itemsASet := map[rune]bool{}

	for _, item := range itemsA {
		itemsASet[item] = true
	}

	for _, item := range itemsB {
		if _, ok := itemsASet[item]; ok {
			commomItems += string(item)
		}
	}

	return commomItems
}

func getPrioritySum(bags []string) int {
	sum := 0

	for _, bag := range bags {
		leftComp := bag[:len(bag)/2]
		rightComp := bag[len(bag)/2:]
		repeatedItem := getCommomItems(leftComp, rightComp)
		sum += getItemPriority(repeatedItem)
	}

	return sum
}

func getGroupCommomItem(group []string) string {
	commomItems := group[0]

	for i := 1; i < len(group); i++ {
		commomItems = getCommomItems(commomItems, group[i])
	}

	return string(commomItems[0])
}

func getGroupPrioritySum(bags []string) int {
	sum := 0
	groupSize := 3
	for i := 0; i <= len(bags)-3; i += groupSize {
		group := bags[i : i+groupSize]

		groupCommomItem := getGroupCommomItem(group)

		sum += getItemPriority(groupCommomItem)
	}

	return sum
}

func part1(input []string) {
	answer := getPrioritySum(input)
	fmt.Println("part 1:", answer)
}

func part2(input []string) {
	answer := getGroupPrioritySum(input)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}

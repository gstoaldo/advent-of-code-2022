package main

import "fmt"

func isFullyContained(pair []int) bool {
	firstContainsSecond := pair[0] <= pair[2] && pair[1] >= pair[3]
	secondContainsFirst := pair[2] <= pair[0] && pair[3] >= pair[1]

	return firstContainsSecond || secondContainsFirst
}

func countFullyContainedPairs(pairs [][]int) int {
	count := 0
	for _, pair := range pairs {
		if isFullyContained(pair) {
			count++
		}
	}

	return count
}

func part1(input [][]int) {
	answer := countFullyContainedPairs(input)
	fmt.Println("part 1:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
}

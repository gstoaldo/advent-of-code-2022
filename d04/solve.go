package main

import (
	"fmt"
	"math"
)

func isFullyContained(pair []int) bool {
	firstContainsSecond := pair[0] <= pair[2] && pair[1] >= pair[3]
	secondContainsFirst := pair[2] <= pair[0] && pair[3] >= pair[1]

	return firstContainsSecond || secondContainsFirst
}

func isOverlap(pair []int) bool {
	maxStart := math.Max(float64(pair[0]), float64(pair[2]))
	minEnd := math.Min(float64(pair[1]), float64(pair[3]))

	return maxStart <= minEnd
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

func countOverlapPairs(pairs [][]int) int {
	count := 0
	for _, pair := range pairs {
		if isOverlap(pair) {
			count++
		}
	}

	return count
}

func part1(input [][]int) {
	answer := countFullyContainedPairs(input)
	fmt.Println("part 1:", answer)
}

func part2(input [][]int) {
	answer := countOverlapPairs(input)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}

package main

import "fmt"

func windowIsUnique(stream inputType, windowSize int) bool {
	if windowSize > len(stream) {
		return false
	}

	window := stream[len(stream)-windowSize:]

	chSet := map[rune]bool{}

	for _, ch := range window {
		if _, ok := chSet[ch]; ok {
			return false
		}

		chSet[ch] = true

	}

	return true
}

func findMarker(stream inputType, windowSize int) int {
	for i := range stream {
		if windowIsUnique(stream[:i], windowSize) {
			return i
		}
	}

	return 0
}

func part1(input inputType) {
	windowSize := 4
	answer := findMarker(input, windowSize)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	windowSize := 14
	answer := findMarker(input, windowSize)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}

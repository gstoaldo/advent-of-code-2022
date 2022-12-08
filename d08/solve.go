package main

import (
	"fmt"
)

type inputType [][]int

func isOnTheEdge(grid inputType, i int, j int) bool {
	return i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1
}

func isInBounds(grid inputType, i int, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func isVisible(grid inputType, i int, j int) bool {
	deltas := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	if isOnTheEdge(grid, i, j) {
		return true
	}

	visible := true

	for _, delta := range deltas {
		ipath, jpath := i+delta[0], j+delta[1]

		for isInBounds(grid, ipath, jpath) {
			visible = grid[ipath][jpath] < grid[i][j]

			if !visible {
				break
			}

			ipath += delta[0]
			jpath += delta[1]
		}

		if visible {
			return true
		}
	}

	return false
}

func countVisibleTrees(grid inputType) int {
	sum := 0

	for i, row := range grid {
		for j := range row {
			if isVisible(grid, i, j) {
				sum++
			}
		}
	}

	return sum
}

func getScenicScore(grid inputType, i int, j int) int {
	deltas := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	if isOnTheEdge(grid, i, j) {
		return 0
	}

	scenicScore := 1

	for _, delta := range deltas {
		directionScore := 0

		ipath, jpath := i+delta[0], j+delta[1]

		for isInBounds(grid, ipath, jpath) {
			directionScore++

			if grid[ipath][jpath] >= grid[i][j] {
				break
			}

			ipath += delta[0]
			jpath += delta[1]
		}

		scenicScore *= directionScore
	}

	return scenicScore
}

func findMaxScenicScore(grid inputType) int {
	max := 0

	for i, row := range grid {
		for j := range row {
			scenicScore := getScenicScore(grid, i, j)

			if scenicScore > max {
				max = scenicScore
			}
		}
	}

	return max
}

func part1(input inputType) {
	answer := countVisibleTrees(input)
	fmt.Println("part 1:", answer)
}

func part2(input inputType) {
	answer := findMaxScenicScore(input)
	fmt.Println("part 2:", answer)
}

func main() {
	input := parseFile("input.txt")

	part1(input)
	part2(input)
}
